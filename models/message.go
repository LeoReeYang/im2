package models

type MsgType int32

const (
	PrivateMessage MsgType = 0
	PublicMessage  MsgType = 1
)

type Message struct {
	ID        string `json:"id"`
	Sender    string `json:"from"`
	Recipient string `json:"to"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}
