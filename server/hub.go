package server

import (
	"time"

	"github.com/fatih/color"

	"github.com/LeoReeYang/im2/connection"
	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/mq"
	"github.com/gorilla/websocket"
)

const (
	CheckPeriod = time.Second * 5
)

type MessageHandler interface {
	Get() *models.Message
	Put(*models.Message) error
}

type ConnectionHandeler interface {
	GetByName(string) *connection.Connection
	Put(*connection.Connection)
	Remove(string)
	GetAll() []string
	CheckConnections()
}

type Hub struct {
	// Register requests from the clients.
	register chan *connection.Connection
	// Leave requests from clients.
	leave chan *connection.Connection

	MessageMeta MessageHandler
	UserMeta    ConnectionHandeler
}

func NewHub(mh MessageHandler, uh ConnectionHandeler) *Hub {
	return &Hub{
		register:    make(chan *connection.Connection, 1024),
		leave:       make(chan *connection.Connection, 1024),
		MessageMeta: mh,
		UserMeta:    uh,
	}
}

func (h *Hub) NewConnection(name string, conn *websocket.Conn) {
	newConnection := connection.NewConnection(name, conn)
	h.RegisterConnection(newConnection)
	newConnection.ListenChannel()
}

func (h *Hub) Transfer(msg *models.Message) {
	receiver := h.UserMeta.GetByName(msg.Recipient)
	receiver.PutSendChannel(msg)
}

func (h *Hub) RegisterConnection(c *connection.Connection) {
	color.Cyan("[ %s ]  connection registering...", c.GetName())
	h.register <- c
}
func (h *Hub) LeaveConnection(c *connection.Connection) {
	color.Cyan("[ %s ]  connection leaving...", c.GetName())
	h.leave <- c
}

func (h *Hub) Run() {
	go h.UserMeta.CheckConnections()

	for {
		select {
		case user := <-h.register:
			h.UserMeta.Put(user)
		case user := <-h.leave:
			h.UserMeta.Remove(user.GetName())
			user.CloseChannel()
		case msg, ok := <-mq.MQ:
			if ok {
				h.Transfer(msg)
			}
		}
	}

}