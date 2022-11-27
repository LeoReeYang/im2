package router

import (
	"flag"
	"net/http"

	v1 "github.com/LeoReeYang/im2/api/v1"
	"github.com/LeoReeYang/im2/docs"
	"github.com/LeoReeYang/im2/global"
	"github.com/LeoReeYang/im2/middlewares"
	"github.com/gorilla/websocket"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func SetupRouters() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	flag.Parse()

	hub := global.Server.GetHub()

	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)

	r.GET("/testchat", v1.ChannelHandle)

	r.GET("/users", func(ctx *gin.Context) {
		data := hub.UserMeta.GetAll()
		ctx.JSON(http.StatusOK, gin.H{
			"users": data,
		})
	})

	protected := r.Group("api/admin")
	{
		protected.Use(middlewares.JWTAuthMiddleware())
		protected.GET("/user", v1.CurrentUser)
	}

	test := r.Group("/test")
	{
		test.GET("/echo", v1.Echo)
	}

	chat := r.Group("/chat")
	{
		chat.Use(middlewares.JWTAuthMiddleware())
		chat.GET("/", v1.ChannelHandle)
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
