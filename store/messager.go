package store

type Messager struct {
	PublicQueue  *MsgQueue
	PrivateQueue *MsgQueue

	Users *UserStore
}

func NewMessager() *Messager {
	return &Messager{
		PublicQueue:  NewMsgQueue(),
		PrivateQueue: NewMsgQueue(),
		Users:        NewUserStore(),
	}
}
