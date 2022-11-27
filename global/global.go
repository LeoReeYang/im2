package global

import (
	"github.com/LeoReeYang/im2/Meta"
	"github.com/LeoReeYang/im2/server"
)

var (
	ConnectionMeta server.ConnectionHandeler = Meta.NewConnectionMeta()
	MessageMeta    server.MessageHandler     = Meta.NewMessageMeta()

	Server *server.Server = server.NewServer(MessageMeta, ConnectionMeta)
)
