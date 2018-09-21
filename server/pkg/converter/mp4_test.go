package converter

import (
	"log"
	"os"
	"testing"
)

func TestCreateMp4(t *testing.T) {
	path, err := CreateMp4("./", "test.png", "saufen.mp3")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
	}
}
