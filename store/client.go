// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
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

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Uuid uint64 = 0

type Client struct {
	id     string
	roomid string
	name   string
	hub    *Hub
	conn   *websocket.Conn
	send   chan *models.Message
	recive chan *models.Message
}

type User struct {
	name   string
	huber  *Huber
	conn   *websocket.Conn
	send   chan *models.Message
	recive chan *models.Message
}

func NewUser(Name string, huber *Huber, conn *websocket.Conn) *User {
	return &User{
		name:   Name,
		huber:  huber,
		conn:   conn,
		send:   make(chan *models.Message, 1024),
		recive: make(chan *models.Message, 1024),
	}
}

// NewClient creates a new client
func NewClient(uid string, roomid string, nickname string, conn *websocket.Conn, hub *Hub) *Client {
	return &Client{id: uid, roomid: roomid, name: nickname, conn: conn, send: make(chan *models.Message, 1024),
		recive: make(chan *models.Message, 1024), hub: hub}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		info := &Info{
			Client: c,
			Room:   c.roomid,
		}
		c.hub.unregister <- info
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg models.Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		c.recive <- &msg
		c.hub.broadcast <- &msg
		fmt.Printf("Client %v get msg from conn:\n msg: %v\n", c.name, msg)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} else {
				err := c.conn.WriteJSON(message)
				if err != nil {
					fmt.Println("Client WriteJson Error: ", err)
					break
				}

				fmt.Printf("Client: < %s > send msg to conn\n msg : %v\n", c.name, message)
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(roomid string, nickname string, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	atomic.AddUint64(&Uuid, 1)

	client := NewClient("test", roomid, nickname, conn, hub)
	info := &Info{
		Client: client,
		Room:   roomid,
	}
	client.hub.register <- info

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func (u *User) read() {
	// defer func() {
	// 	u.huber.unregister <- u
	// }()
	u.conn.SetReadLimit(maxMessageSize)
	u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error { u.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg models.Message
		err := u.conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		u.huber.mq <- &msg
		color.Magenta("Client %v get msg from conn:\n msg: %v\n", u.name, msg)
	}
}

func (u *User) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.conn.Close()
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
					break
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

func Handle(w http.ResponseWriter, r *http.Request, huber *Huber, name string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewUser(name, huber, conn)
	client.huber.register <- client

	go client.read()
	go client.write()
}
