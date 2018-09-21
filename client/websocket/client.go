package client

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func ListenAndWrite(addr string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := NewConn(addr)
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		if err := Listen(c); err != nil {
			return
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			if err := Write(c, t.String()); err != nil {
				return
			}
		case <-interrupt:
			disconnectGracefully(c, done)
			return
		}
	}
}

func NewConn(addr string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}

func Listen(c *websocket.Conn) error {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return err
		}
		log.Printf("recv: %s", message)
	}
}

func Write(c *websocket.Conn, msg string) error {
	err := c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
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