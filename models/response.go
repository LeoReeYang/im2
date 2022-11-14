package models

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewStatus(code int, msg string) *Status {
	return &Status{
		Code: code,
		Msg:  msg,
	}
}

type GetMessageResponse struct {
	Status   *Status    `json:"status"`
	Messages []*Message `json:"messages"`
}

type SendMessageResponse struct {
	Status *Status `json:"status"`
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type LoginRsp struct {
	Message string `json:"message"`
}
type RegisterRsp struct {
	Message string `json:"message"`
}
