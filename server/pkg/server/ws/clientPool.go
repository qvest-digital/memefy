package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var clientMap = make(map[string]*websocket.Conn)

var mutex = sync.RWMutex{}

func AddClient(id string, c *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	clientMap[id] = c
}

func RemoveClient(id string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clientMap, id)
}

func GetClient(id string) *websocket.Conn {
	mutex.RLock()
	defer mutex.RUnlock()
	return clientMap[id]
}

func TriggerMeme(file string) error {
	mutex.RLock()
	defer mutex.RUnlock()
	msg, err := websocket.NewPreparedMessage(websocket.TextMessage, []byte(file))
	if err != nil {
		return err
	}
	for k := range clientMap {
		if err := clientMap[k].WritePreparedMessage(msg); err != nil {
			log.Printf("Could not trigger meme for client %s: %s", k, err.Error())
		}
	}
	return nil
}


