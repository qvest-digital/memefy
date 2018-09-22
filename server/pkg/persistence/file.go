package persistence

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/birkirb/loggers.v1/log"
)

//returns the filename of the saved file or an error if it occurs
func SaveMultipartFile(r *http.Request, partname string, storagePath string) (string, error) {
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
	if err != nil {
		return "", err
	}
	defer f.Close()

	//save our file to our path
	written, err := io.Copy(f, file)
	if err != nil {
		return "", err
	}

	log.Infof("File '%s' saved as '%s', '%d' bytes", handler.Filename, savepath, written)
	return savepath, nil
}

func SaveMetaData(r *http.Request, storagePath string) (string, error) {
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

func GetMeme(name, storagePath string) ([]byte, error) {
	return ioutil.ReadFile(storagePath + name + "/out.mp4")
}
