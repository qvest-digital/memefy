package play

import (
	"log"
	"testing"
)

func TestPlayMeme(t *testing.T) {
	if err := PlayMeme("out.mp4"); err != nil {
		log.Fatal(err)
	}
}
