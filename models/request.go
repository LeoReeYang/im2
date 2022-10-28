package models

type GetMessagesRequest struct {
	Msg Message `json:"msg"`
}

type SendMessagesRequest struct {
	Msg Message `json:"msg"`
}
