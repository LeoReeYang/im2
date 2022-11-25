package store

import (
	"sync"

	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
)

type Huber struct {
	// Unregistered clients.
	unregister chan *User
	// Register requests from the clients.
	register chan *User
	// Inbound messages from the clients.
	mq chan *models.Message
	// all clients
	clients map[string]*User
	// check whether User online
	Clients map[*User]bool

	Messager MessageStore
	locker   sync.RWMutex
}

func NewHuber(ms MessageStore) *Huber {
	return &Huber{
		unregister: make(chan *User),
		register:   make(chan *User),
		mq:         make(chan *models.Message, 1024),
		clients:    make(map[string]*User),
		Clients:    make(map[*User]bool),
		Messager:   ms,
	}
}

func (h *Huber) Run() {
	for {
		select {
		case user := <-h.register:
			h.registe(user)
		case user := <-h.unregister:
			h.unregiste(user)
		case msg := <-h.mq:
			h.transfer(msg)
		}
	}
}

func (h *Huber) registe(u *User) {
	color.Cyan("%s connect to huber", u.name)

	h.locker.Lock()
	h.clients[u.name] = u
	h.locker.Unlock()
}

func (h *Huber) unregiste(u *User) {
	color.Cyan("%s disconnect...", u.name)

	h.locker.Lock()
	delete(h.clients, u.name)
	delete(h.Clients, u)
	h.locker.Unlock()

	u.conn.Close()
}

func (h *Huber) messageStore(msg *models.Message) error {
	return h.Messager.Put(msg)
}

func (h *Huber) transfer(msg *models.Message) {
	h.locker.RLock()
	if user, ok := h.clients[msg.Recipient]; ok {
		color.Yellow("%s 's channel receive a msg", msg.Recipient)
		user.send <- msg
	}
	h.locker.RUnlock()

	err := h.messageStore(msg)
	if err != nil {
		color.Yellow("internl faild")
	}
}

func (h *Huber) LeaveEnqueue(u *User) {
	h.unregister <- u
}

func (h *Huber) EnterEnqueue(u *User) {
	h.register <- u
}

func (h *Huber) MessageEnqueue(msg *models.Message) {
	h.mq <- msg
}

func (h *Huber) GetAllUsers() []string {
	ret := []string{}
	h.locker.RLock()
	for name := range h.clients {
		ret = append(ret, name)
	}
	h.locker.RUnlock()

	return ret
}
