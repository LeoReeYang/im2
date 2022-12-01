package router

import (
	"flag"
	"net/http"

	v1 "github.com/LeoReeYang/im2/api/v1"
	"github.com/LeoReeYang/im2/docs"
	"github.com/LeoReeYang/im2/global"
	"github.com/LeoReeYang/im2/middlewares"
	"github.com/LeoReeYang/im2/utils/token"
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

	r.GET("/messages/:name", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		name := ctx.Param("name")
		_, parsed_name, err := token.ExtractTokenInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusExpectationFailed, gin.H{
				"message": err,
			})
			return
		}

		if name != parsed_name {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "can't request others messages.",
			})
			return
		}
		data := hub.MessageMeta.Get(name)
		ctx.JSON(http.StatusOK, gin.H{
			"Messages": data,
		})
	})

	protected := r.Group("api/online", middlewares.JWTAuthMiddleware())
	{
		protected.GET("/me", v1.Me)
		protected.GET("/all", v1.All)
	}

	test := r.Group("/test")
	{
		test.GET("/echo", v1.Echo)
		test.GET("/chat", v1.ChannelHandle)
	}

	chat := r.Group("/chat")
	{
		chat.Use(middlewares.JWTAuthMiddleware())
		chat.GET("/", v1.ChannelHandle)
		chat.GET("/users", func(ctx *gin.Context) {
			data := hub.UserMeta.All()
			ctx.JSON(http.StatusOK, gin.H{
				"users": data,
			})
		})
	}

	users := r.Group("/user", middlewares.JWTAuthMiddleware())
	{
		users.GET("", v1.Query)
		users.DELETE("", v1.Delete)
		users.PUT("", v1.Update)
		users.POST("/friend", v1.AddFriend)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(*addr)
}
