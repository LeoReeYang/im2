package v1

import (
	"log"

	"github.com/LeoReeYang/im2/global"
	"github.com/gin-gonic/gin"
)

func ChannelHandle(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	name := ctx.Query("nickname")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	hub := global.Server.GetHub()
	hub.NewConnection(name, conn)
}
