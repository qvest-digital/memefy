package main

import (
	"log"
	ws "memefy/client/websocket"
	"os"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	server := os.Getenv("MEMEFY_SERVER")
	if server == "" {
		server = "gomano.de:8080"
	}
	mid, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	ws.ListenAndWrite(server, "/client/"+mid)
}
