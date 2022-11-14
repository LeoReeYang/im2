package models

type GetMessagesRequest struct {
	Msg Message `json:"msg"`
}

type SendMessagesRequest struct {
	Seq string  `json:"seq"`
	Cmd string  `json:"cmd"`
	Msg Message `json:"msg"`
}

type Request struct {
	Status int
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
