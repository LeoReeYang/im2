package store

import (
	"sync"
	"time"

	"github.com/LeoReeYang/im2/models"
	"github.com/gorilla/websocket"
)

type status bool

const (
	OnLine  status = true
	OffLine status = false
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

var Uuid uint64 = 0

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

type Meta struct {
	// Registered clients.
	ClientsStatus map[*User]status

	// Clients List
	ClientsList map[string]*User

	// Register requests from the clients.
	Register chan *User

	// Unregister requests from clients.
	Unregister chan *User

	// locker sync.RWMutex
	sync.RWMutex
}

func NewMeta() *Meta {
	return &Meta{
		ClientsStatus: make(map[*User]status),
		ClientsList:   make(map[string]*User),
		Register:      make(chan *User),
		Unregister:    make(chan *User),
	}
}
