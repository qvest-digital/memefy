package main

import ws "memefy/client/websocket"

func main() {
	ws.ListenAndWrite("localhost:8080", "/client/123")
}
