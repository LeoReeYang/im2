package store

import (
	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
)

type Hub struct {
	// mq for Messages.
	mq chan *models.Message
	// UserHandel handles Users who connecting to the Server.
	UserHandel UserHandeler
	// Messager store the Message and deal transactions with Database.
	Messager MessageStore
}

func NewHub(ms MessageStore) *Hub {
	return &Hub{
		mq:         make(chan *models.Message, 1024),
		UserHandel: NewUserMeta(),
		Messager:   ms,
	}
}

func (h *Hub) Run() {
	go h.UserHandel.Run()
	for {
		msg := <-h.mq
		h.UserHandel.Transfer(msg)

		err := h.StoreMessage(msg)
		if err != nil {
			color.Yellow("internl message store faild.")
		}
	}
}

func (h *Hub) StoreMessage(msg *models.Message) error {
	return h.Messager.Put(msg)
}

func (h *Hub) LeaveEnqueue(u *User) {
	h.UserHandel.PutUserToLeaveChannel(u)
}

func (h *Hub) EnterEnqueue(u *User) {
	h.UserHandel.PutUserToRegisterChannel(u)
}

func (h *Hub) MessageEnqueue(msg *models.Message) {
	h.mq <- msg
}
