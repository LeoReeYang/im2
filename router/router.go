package router

import (
	_ "github.com/LeoReeYang/im2/api"
	// _ "github.com/LeoReeYang/im2/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouters() {
	r := gin.Default()

	r.GET("/echo", func(ctx *gin.Context) {
		// api.v1.echo(ctx.Writer, ctx.Request)
	})
}
