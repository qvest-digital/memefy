package main

import (
	"log"
	ws "memefy/client/websocket"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	mid, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	ws.ListenAndWrite("localhost:8080", "/client/"+mid)
}
