package store

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (u *User) read() {
	defer func() {
		u.hub.UserHandel.PutUserToLeaveChannel(u)
	}()
	u.conn.SetReadLimit(maxMessageSize)
	u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error { u.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg models.Message
		err := u.conn.ReadJSON(&msg)
		if err != nil {
			color.Red("Huber User ReadJSON:", err)
			break
		}
		u.hub.MessageEnqueue(&msg)
		color.Magenta("Client %v get msg from conn:\n msg: %v\n", u.name, msg)
	}
}

func (u *User) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.hub.UserHandel.PutUserToLeaveChannel(u)
	}()

	for {
		select {
		case message, ok := <-u.send:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} else {
				err := u.conn.WriteJSON(message)
				if err != nil {
					fmt.Println("Client WriteJson Error: ", err)
					return
				}

				color.Magenta("Client: < %s > send msg to conn\nmsg : %v\n", u.name, message)
			}
		case <-ticker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func Handle(w http.ResponseWriter, r *http.Request, hub *Hub, name string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	user := NewUser(name, hub, conn)
	user.hub.UserHandel.PutUserToRegisterChannel(user)

	go user.read()
	go user.write()
}
