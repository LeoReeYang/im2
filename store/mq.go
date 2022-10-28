package store

import "im/models"

type MsgQueue struct {
	msgqueue chan *models.Message
}

func (mq *MsgQueue) NewMsgQueue() *MsgQueue {
	return &MsgQueue{
		msgqueue: make(chan *models.Message, 100),
	}
}

func (mq *MsgQueue) PutMsg(msg *models.Message) error {
	mq.msgqueue <- msg
	return nil
}

func (mq *MsgQueue) GetMsg() *models.Message {
	msg := <-mq.msgqueue
	return msg
}
