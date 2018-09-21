package main

import (
	"log"
	"memefy/client/play"
)

type Player func(string) error

var player Player = play.PlayMeme

func main() {
	if err := player("small.mp4"); err != nil {
		log.Fatal(err)
	}
}
