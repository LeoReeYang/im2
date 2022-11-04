package main

import (
	"flag"

	"github.com/LeoReeYang/im2/server"
	"github.com/LeoReeYang/im2/store"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	// router.SetupRouters()
	flag.Parse()

	r := gin.Default()

	server := server.NewServer(store.NewMsgQueue())

	hub := server.GetHub()
	go hub.Run()

	r.GET("/echo", func(ctx *gin.Context) {
		// api.v1.echo(ctx.Writer, ctx.Request)
		// api.echo(ctx.Writer, ctx.Request)
		// echo(ctx)
		// api.echo
		// v1.echo(ctx)
	})

	r.GET("/chat", func(ctx *gin.Context) {
		roomid := ctx.DefaultQuery("roomid", "001")
		nickName := ctx.DefaultQuery("nickname", "unknown user")

		store.ServeWs(roomid, nickName, hub, ctx.Writer, ctx.Request)
	})

	r.Run(*addr)
}
