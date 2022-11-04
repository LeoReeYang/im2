package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/LeoReeYang/im2/models"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func TestClient(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	msg := &models.Message{
		ID:        "123",
		Sender:    123,
		Recipient: "server",
		Type:      "message",
		Content:   "test",
		Timestamp: time.Now().Unix(),
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	b, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", string(b))
	c.WriteMessage(websocket.BinaryMessage, b)
}
