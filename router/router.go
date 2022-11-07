package router

import (
	"flag"

	v1 "github.com/LeoReeYang/im2/api/v1"
	"github.com/LeoReeYang/im2/docs"
	"github.com/LeoReeYang/im2/server"
	"github.com/LeoReeYang/im2/store"

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

	r.GET("/ws/:nickname", func(ctx *gin.Context) {
		w := ctx.Writer
		r := ctx.Request
		nickname := ctx.Param("nickname")
		store.ServeWs("001", nickname, hub, w, r)
	})

	test := r.Group("/test")
	{
		test.GET("/echo", v1.Echo)
	}

	users := r.Group("/user")
	{
		users.GET("/:id", v1.QueryUser)
		users.DELETE("/:id", v1.DeleteUser)
		users.PUT("/:id", v1.UpdateUser)
		users.POST("/", v1.CreateUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(*addr)
}
