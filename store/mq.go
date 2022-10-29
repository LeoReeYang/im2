package store

import "im/models"

// import "im/models"

type MsgQueue struct {
	msgqueue chan *models.Message
}

func NewMsgQueue() *MsgQueue {
	return &MsgQueue{
		msgqueue: make(chan *models.Message, 100),
	}
}

func (mq *MsgQueue) Put(msg *models.Message) error {
	mq.msgqueue <- msg
	return nil
}

func (mq *MsgQueue) Get() *models.Message {
	msg := <-mq.msgqueue
	return msg
}
