package models

import "time"

type MsgType int32

const (
	PrivateMessage MsgType = 0
	PublicMessage  MsgType = 1
)

type Message struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}
