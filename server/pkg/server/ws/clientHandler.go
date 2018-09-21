package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: allOrigins} // use default options

// MemeDiffer should check for missing memes on the client side
type MemeDiffer func(oldMemes, currentMemes []string) []string

// MemeLister should return the current meme selection
type MemeLister func() []string

func NewMemeHandleFunc(memeDiffer MemeDiffer, memeLister MemeLister) http.HandlerFunc {
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

			memeDiff := memeDiffer(clientMsg.CurrentMemes, memeLister())
			c.WriteJSON(&NewMemes{memeDiff})
			if err != nil {
				log.Printf("Could not sent new meme listing to client %s: %s", id, err.Error())
			}
		}
	}
}

func allOrigins(r *http.Request) bool {
	return true
}
