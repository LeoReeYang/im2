package mq

import "github.com/LeoReeYang/im2/models"

// var Mq *SimpleMsgQueue = NewQueue()

var MQ chan *models.Message = make(chan *models.Message, 1024)

type MessageQueue interface {
	Push(*models.Message)
	Get()
}

// type SimpleMsgQueue struct {
// 	Mq chan *models.Message
// }

// func NewQueue() *SimpleMsgQueue {
// 	return &SimpleMsgQueue{
// 		Mq: make(chan *models.Message, 1024),
// 	}
// }

// func (s *SimpleMsgQueue) Push(msg *models.Message) {
// 	s.Mq <- msg
// }

// func (s *SimpleMsgQueue) Get() (msg *models.Message) {
// 	msg = <-s.Mq
// 	return
// }
