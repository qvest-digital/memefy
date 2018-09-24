package client

import (
	"encoding/json"
	"fmt"
	"log"
	"memefy/client/persistence"
	"memefy/client/play"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func ListenAndWrite(addr, path string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := newConn(addr, path)
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		if err := Listen(c); err != nil {
			return
		}
	}()

	go func() {
		for {
			currentMemes, err := persistence.ListMemes()
			if err != nil {
				log.Fatalf("Could not list current memes: %s", err.Error())
			}
			c.WriteJSON(&ClientSyncRequest{CurrentMemes: currentMemes})
			<-time.After(1 * time.Second)
		}

	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			disconnectGracefully(c, done)
			return
		}
	}
}

func newConn(addr, path string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: addr, Path: path}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}

func Listen(c *websocket.Conn) error {
	for {
		msgType, msgReader, err := c.NextReader()
		if err != nil {
			log.Println("recv error: ", err)
			return err
		}
		if msgType == websocket.BinaryMessage {
			nameLenByte := make([]byte, 1)
			n, err := msgReader.Read(nameLenByte)
			if err != nil {
				log.Println("binary read error: ", err)
				return err
			}
			if n != 1 {
				log.Println("binary read error, read more then one length byte")
				return fmt.Errorf("Instead of reading 1 length byte %d were read", n)
			}
			nameLen := int(nameLenByte[0])
			nameBytes := make([]byte, nameLen)
			n, err = msgReader.Read(nameBytes)
			if err != nil {
				log.Println("binary read error while reading name: ", err)
				return err
			}
			if n != nameLen {
				log.Println("binary read error, read more then one length byte")
				return fmt.Errorf("Instead of reading %d length byte %d were read", nameLen, n)
			}
			log.Println("Receiving meme: " + string(nameBytes))
			persistence.SaveMeme(string(nameBytes), msgReader)
		} else {
			trigger := &Trigger{}
			err := json.NewDecoder(msgReader).Decode(trigger)
			if err != nil {
				log.Println("read error: ", err)
				return err
			}
			log.Println("got ", trigger)
			err = play.PlayMeme(trigger.Meme)
			if err != nil {
				log.Printf("Error: %s", err.Error())
			}
		}
	}
}

func disconnectGracefully(c *websocket.Conn, done chan struct{}) {
	log.Println("interrupt")

	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
}
