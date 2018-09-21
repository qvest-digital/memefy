package ws

import (
	"log"
	"memefy/server/pkg/persistence"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: allOrigins} // use default options

// MemeDiffer should check for missing memes on the client side
type MemeDiffer func(oldMemes, currentMemes []string) []string

// MemeLister should return the current meme selection
type MemeLister func() []string

func WebSocketClientHandler(memeDiffer MemeDiffer, memeLister MemeLister, storagePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		id := mux.Vars(r)["clientId"]
		defer RemoveClient(id)
		AddClient(id, c)

		for {
			clientMsg := &ClientRegistration{}
			err := c.ReadJSON(clientMsg)
			if err != nil {
				log.Printf("Client %s sent unparseable message: %s", id, err.Error())
			}

			memeDiff := memeDiffer(memeLister(), clientMsg.CurrentMemes)
			for i := range memeDiff {
				log.Println("Sending meme to client " + memeDiff[i])
				meme, err := persistence.GetMeme(memeDiff[i], storagePath)
				if err != nil {
					log.Printf("Could not send meme %s to client %s: %s", memeDiff[i], id, err.Error())
				}
				byteName := []byte(memeDiff[i])
				nameLen := byte(len(byteName))

				byteMsg := []byte{}
				byteMsg = append(byteMsg, nameLen)
				byteMsg = append(byteMsg, byteName...)
				byteMsg = append(byteMsg, meme...)
				log.Printf("Writing %d bytes to client %s", len(byteMsg), id)
				err = c.WriteMessage(websocket.BinaryMessage, byteMsg)
				if err != nil {
					log.Printf("Could not send binary meme %s to client %s: %s", memeDiff[i], id, err.Error())
				}
			}

			if err != nil {
				log.Printf("Could not sent new meme listing to client %s: %s", id, err.Error())
			}
		}
	}
}

func allOrigins(r *http.Request) bool {
	return true
}

func NewFsMemeLister(basePath string) MemeLister {
	return func() []string {
		storageDir, err := os.Open(basePath)
		defer storageDir.Close()
		if err != nil {
			log.Printf("Failed opening directory: %s", err)
		}

		memeList := []string{}
		list, _ := storageDir.Readdir(0)
		for _, f := range list {
			if f.IsDir() {
				memeList = append(memeList, f.Name())
			}
		}

		return memeList
	}
}

func NewMemeDiffer() MemeDiffer {
	return func(a, b []string) []string {
		m := make(map[string]bool)
		diff := []string{}

		for _, item := range b {
			m[item] = true
		}

		for _, item := range a {
			if _, ok := m[item]; !ok {
				diff = append(diff, item)
			}
		}
		return diff
	}
}
