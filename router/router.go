package router

import (
	"flag"

	v1 "github.com/LeoReeYang/im2/api/v1"
	"github.com/LeoReeYang/im2/server"
	"github.com/LeoReeYang/im2/store"

	docs "github.com/LeoReeYang/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func SetupRouters() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	flag.Parse()

	server := server.NewServer(store.NewMsgQueue())

	hub := server.GetHub()
	go hub.Run()

	r.GET("/echo", v1.Echo)
	r.GET("/char", func(ctx *gin.Context) {
		v1.ServerWs(ctx, hub)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(*addr)
}
