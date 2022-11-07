// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"fmt"

	"github.com/LeoReeYang/im2/models"
)

type Info struct {
	Client *Client
	Room   string
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	rooms map[string]map[*Client]bool
	//Unregistered clients.
	unregister chan *Info
	// Register requests from the clients.
	register chan *Info
	// Inbound messages from the clients.
	broadcast chan *models.Message
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*Client]bool),
		unregister: make(chan *Info),
		register:   make(chan *Info),
		broadcast:  make(chan *models.Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// Register a client.
		case info := <-h.register:
			h.RegisterNewClient(info)
		// Unregister a client.
		case info := <-h.unregister:
			h.RemoveClient(info)
		// Broadcast a message to all clients.
		case message := <-h.broadcast:
			//Check if the message is a type of "message"
			h.HandleMessage(message)
		}
	}
}

// function check if room exists and if not create it and add client to it
func (h *Hub) RegisterNewClient(info *Info) {
	// id refer to different rooms
	roomid := info.Room
	client := info.Client

	connections := h.rooms[roomid]
	if connections == nil {
		connections = make(map[*Client]bool)
		h.rooms[roomid] = connections
	}
	h.rooms[roomid][client] = true

	fmt.Printf("Size of clients %d ,of room %s: \n", len(h.rooms[client.roomid]), roomid)
}

// function to remvoe client from room
func (h *Hub) RemoveClient(info *Info) {
	roomid := info.Room
	client := info.Client

	if _, ok := h.rooms[roomid]; ok {
		delete(h.rooms[roomid], client)
		close(client.send)
		fmt.Printf("Removed client <%v> from room: '%v'\n", client, roomid)
	}
}

// function to handle message based on type of message
func (h *Hub) HandleMessage(message *models.Message) {
	// fmt.Printf("hub handle msg : %v\n", message)

	// if message.Type != "message" {
	// 	log.Fatal("not equal!\n")
	// }

	switch message.Type {
	case "message":
		clients := h.rooms[message.ID]
		for client := range clients {
			select {
			case client.send <- message:
				// fmt.Println("msg -> client send channel ok!")
				fmt.Printf("broadcast channel handle msg from %v to %v \n", message.Sender, message.Recipient)
			default:
				close(client.send)
				delete(h.rooms[message.ID], client)
			}
		}
	case "notification":
		fmt.Println("Notification: ", message.Content)
		clients := h.rooms[message.Recipient]
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.rooms[message.Recipient], client)
			}
		}
	}
}
