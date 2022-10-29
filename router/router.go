package router

import (
	"im/server"
	"im/store"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouters() {
	r := gin.Default()
	// r.LoadHTMLFiles("im/home.html")
	r.LoadHTMLGlob("assets/*")
	s := server.NewServer()
	s.Setup()

	r.GET("/ws/:roomId", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")
		store.ServeWs(roomId, s.GetHub(), ctx.Writer, ctx.Request)
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	err := r.Run("localhost:12312")
	if err != nil {
		log.Fatal(err)
	}
}
