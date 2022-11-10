package store

import (
	"sync"

	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
)

type Huber struct {
	//Unregistered clients.
	unregister chan *User
	// Register requests from the clients.
	register chan *User
	// Inbound messages from the clients.
	mq chan *models.Message
	// all clients
	clients map[string]*User

	locker sync.RWMutex
}

func NewHuber() *Huber {
	return &Huber{
		unregister: make(chan *User),
		register:   make(chan *User),
		mq:         make(chan *models.Message),
		clients:    make(map[string]*User),
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
	go func() {
		color.Cyan("%s connect to huber", u.name)
		h.locker.Lock()
		h.clients[u.name] = u
		h.locker.Unlock()
	}()
}

func (h *Huber) unregiste(u *User) {
	go func() {
		color.Cyan("%s disconnect", u.name)
		h.locker.Lock()
		delete(h.clients, u.name)
		h.locker.Unlock()

		u.conn.Close()
		// close(u.send)
		// close(u.recive)
	}()

}

func (h *Huber) transfer(msg *models.Message) {
	go func() {
		h.locker.RLock()
		for name, user := range h.clients {
			if name == msg.Recipient {
				color.Yellow("%s 's channel get a msg", name)
				user.send <- msg
				break
			}
		}
		h.locker.RUnlock()
	}()
}
