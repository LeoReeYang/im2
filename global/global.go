package global

import (
	"github.com/LeoReeYang/im2/Meta"
	"github.com/LeoReeYang/im2/server"
)

var (
	ConnectionMeta server.ConnectionHandeler
	MessageMeta    server.MessageHandler
	Server         *server.Server
)

func Init() {
	ConnectionMeta = Meta.NewConnectionMeta()
	MessageMeta = Meta.NewMessageMeta()
	Server = server.NewServer(MessageMeta, ConnectionMeta)
}
