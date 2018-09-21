package persistence

import (
	"io/ioutil"
)

const basePath = "files/"

func ListMemes() (memes []string, err error) {
	contents, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	for i := range contents {
		if contents[i].IsDir() {
			memes = append(memes, contents[i].Name())
		}
	}
	return
}
