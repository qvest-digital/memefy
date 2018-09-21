package server

import (
	"encoding/json"
	"fmt"
	"io"
	"memefy/server/pkg/config"
	"memefy/server/pkg/server/ws"
	"memefy/server/pkg/util"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type AdminHandler struct {
	cfg *config.Config
}

type linksInfo struct {
	Links map[string]link `json:"_links"`
}

type link struct {
	Href   string   `json:"href"`
	Method []string `json:"method"`
}

//IndexHandler returns a service self description
func (h *AdminHandler) IndexHandler(router *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		links := make(map[string]link, 20)

		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			name := route.GetName()
			href, _ := route.GetPathTemplate()
			methods, _ := route.GetMethods()

			links[name] = link{
				Href:   href,
				Method: methods,
			}

			return nil
		})

		linksInfos := linksInfo{
			Links: links,
		}

		js, err := json.Marshal(linksInfos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

func (h *AdminHandler) AdminInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminInfoTemplate := `{
			"build": {
				"BuildNumber": "%s",
				"CommitShortHash": "%s",
				"BuildTime": "%s"
			}
		}`
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, adminInfoTemplate, config.BuildNumber, config.CommitShortHash, config.BuildTime)
	}
}

// HealthCheckHandler checks if the service can handle traffic and returns 200 or 503
func (h *AdminHandler) HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// -----

func (h *AdminHandler) PlayMemeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		memelist := ws.NewFsMemeLister(h.cfg.StoragePath)()
		if !util.Contains(memelist, name) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := ws.TriggerMeme(name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func (h *AdminHandler) PostMemeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//parse a request body as multipart/form-data
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		name := r.FormValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("No name given")
			return
		}

		path := h.cfg.StoragePath + "/" + name
		err = os.MkdirAll(path, 0777)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = saveMultipartFile(r, "pic", path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = saveMultipartFile(r, "sound", path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = saveMetaData(r, path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonContent, _ := json.Marshal(map[string]interface{}{
			"name":  name,
			"pic":   fiileEndpoint + name + "/pic",
			"sound": fiileEndpoint + name + "/sound",
			"meta":  fiileEndpoint + name + "/meta",
		})

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonContent)
	}
}

//returns the filename of the saved file or an error if it occurs
func saveMultipartFile(r *http.Request, partname string, storagePath string) (string, error) {
	//retrieve the file from form data
	file, handler, err := r.FormFile(partname)
	defer file.Close()
	if err != nil {
		return "", err
	}

	//this is path which we want to store the file
	//	savepath := storagePath + "/" + handler.Filename
	savepath := storagePath + "/" + partname
	f, err := os.OpenFile(savepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return "", err
	}

	//save our file to our path
	written, err := io.Copy(f, file)
	if err != nil {
		return "", err
	}

	log.Infof("File '%s' saved as '%s', '%d' bytes", handler.Filename, savepath, written)
	return savepath, nil
}

func saveMetaData(r *http.Request, storagePath string) (string, error) {
	content := r.FormValue("meta")

	//this is path which we want to store the file
	savepath := storagePath + "/meta"
	f, err := os.OpenFile(savepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return "", err
	}

	//save our meta to our path
	written, err := f.WriteString(content)
	if err != nil {
		return "", err
	}

	log.Infof("Metadata saved as '%s', '%d' bytes", savepath, written)
	return savepath, nil
}
