package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWEBs(t *testing.T) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		for {
			msgtype, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s receive: %s\n", conn.RemoteAddr(), string(msg))
			if err = conn.WriteMessage(msgtype, msg); err != nil {
				return
			}
		}
	}
	http.HandleFunc("/websocket", handler)

	http.ListenAndServe(":12345", http.HandlerFunc(handler))
}

func TestName(t *testing.T) {

}
