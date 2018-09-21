package server

import (
	"encoding/json"
	"fmt"
	"memefy/server/pkg/config"
	"memefy/server/pkg/converter"
	"memefy/server/pkg/persistence"
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

type meme struct {
	Name    string `json:"name"`
	Picture string `json:"pic"`
	Sound   string `json:"sound"`
	Meta    string `json:"meta"`
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

// -----

func (h *AdminHandler) GetMemeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memelist := ws.NewFsMemeLister(h.cfg.StoragePath)()
		result := make([]*meme, len(memelist))

		for i, memeName := range memelist {
			result[i] = &meme{
				Name:    memeName,
				Picture: fileEndpoint + memeName + "/pic",
				Sound:   fileEndpoint + memeName + "/sound",
				Meta:    fileEndpoint + memeName + "/meta",
			}
		}

		rawResult, err := json.Marshal(result)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(rawResult)
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

		_, err = persistence.SaveMultipartFile(r, "pic", path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = persistence.SaveMultipartFile(r, "sound", path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = persistence.SaveMetaData(r, path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonContent, _ := json.Marshal(&meme{
			Name:    name,
			Picture: fileEndpoint + name + "/pic",
			Sound:   fileEndpoint + name + "/sound",
			Meta:    fileEndpoint + name + "/meta",
		})

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonContent)
		converter.CreateMp4(h.cfg.StoragePath+name+"/", "pic", "sound")
	}
}
