package store

type Messager struct {
	PublicQueue  MsgQueue
	PrivateQueue MsgQueue

	Users UserStore
}

func (m *Messager) NewMessager() *Messager {
	return &Messager{
		PublicQueue:  *m.PublicQueue.NewMsgQueue(),
		PrivateQueue: *m.PrivateQueue.NewMsgQueue(),
		Users:        *m.Users.NewUserStore(),
	}
}
