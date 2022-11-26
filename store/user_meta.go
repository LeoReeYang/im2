package store

import (
	"sync"

	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

type status bool

const (
	OnLine  status = true
	OffLine status = false
)

type User struct {
	name   string
	hub    *Hub
	conn   *websocket.Conn
	send   chan *models.Message
	recive chan *models.Message
}

func NewUser(Name string, hub *Hub, conn *websocket.Conn) *User {
	return &User{
		name:   Name,
		hub:    hub,
		conn:   conn,
		send:   make(chan *models.Message, 1024),
		recive: make(chan *models.Message, 1024),
	}
}

type UserMeta struct {
	// users
	Users map[string]*User

	// Register requests from the clients.
	Register chan *User

	// Leave requests from clients.
	Leave chan *User

	// locker sync.RWMutex
	sync.RWMutex
}

func (um *UserMeta) Run() {
	for {
		select {
		case user := <-um.Register:
			um.registe(user)
		case user := <-um.Leave:
			um.unregiste(user)
		}
	}
}

func (um *UserMeta) Transfer(msg *models.Message) {
	um.RLock()
	if user, ok := um.Users[msg.Recipient]; ok {
		color.Yellow("%s 's channel receive a msg", msg.Recipient)
		user.send <- msg
	}
	um.RUnlock()
}

func NewUserMeta() *UserMeta {
	return &UserMeta{
		Users:    make(map[string]*User),
		Register: make(chan *User),
		Leave:    make(chan *User),
	}
}

func (um *UserMeta) registe(u *User) {
	color.Cyan("%s connect to huber", u.name)
	um.Lock()
	um.Users[u.name] = u
	um.Unlock()
}

func (um *UserMeta) unregiste(u *User) {
	color.Cyan("%s disconnect...", u.name)

	um.Lock()
	delete(um.Users, u.name)
	um.Unlock()

	u.conn.Close()
}

func (um *UserMeta) PutUserToRegisterChannel(u *User) {
	um.Register <- u
}

func (um *UserMeta) PutUserToLeaveChannel(u *User) {
	um.Leave <- u
}

func (um *UserMeta) GetAllUsers() []string {
	users := []string{}

	um.RLock()
	for username := range um.Users {
		users = append(users, username)
	}
	um.RUnlock()

	return users
}
