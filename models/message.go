package models

import "time"

type MsgType int32

const (
	PrivateMessage MsgType = 0
	PublicMessage  MsgType = 1
)

type Message struct {
	From      uint64    `json:"from"`
	To        uint64    `json:"to"`
	MsgType   int32     `json:"msgType"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
