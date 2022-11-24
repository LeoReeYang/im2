package router

import (
	"flag"
	"net/http"

	v1 "github.com/LeoReeYang/im2/api/v1"
	"github.com/LeoReeYang/im2/docs"
	"github.com/LeoReeYang/im2/middlewares"
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

	server := server.NewServer(store.NewSimpleStore())

	hub := server.GetHub()
	go hub.Run()

	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)

	r.GET("/testchat", func(ctx *gin.Context) {
		w := ctx.Writer
		r := ctx.Request

		name := ctx.Query("nickname")
		// color.Magenta("nickname", name)
		store.Handle(w, r, hub, name)
	})

	r.GET("/users", func(ctx *gin.Context) {
		data := hub.GetAllUsers()
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
		chat.GET("/", func(ctx *gin.Context) {
			w := ctx.Writer
			r := ctx.Request

			name := ctx.Query("nickname")
			// color.Magenta("nickname", name)
			store.Handle(w, r, hub, name)
		})
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
