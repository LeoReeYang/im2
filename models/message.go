package models

import "time"

type MsgType int32

const (
	PrivateMessage MsgType = 0
	PublicMessage  MsgType = 1
)

type Message struct {
	ID        string    `json:"id"`
	Sender    uint64    `json:"from"`
	Recipient string    `json:"to"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
