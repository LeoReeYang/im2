package v1

import "github.com/gin-gonic/gin"

func echo(ctx *gin.Context) {
	ctx.String(200, "123")
}
