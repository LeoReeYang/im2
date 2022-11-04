package v1

import (
	"net/http"

	"github.com/LeoReeYang/im2/store"
	"github.com/gin-gonic/gin"
)

// @BasePath

// PingExample godoc
// @Summary echo example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /echo [get]
func Echo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "helloworld")
}

func ServerWs(ctx *gin.Context, hub *store.Hub) {
	roomid := ctx.DefaultQuery("roomid", "001")
	nickName := ctx.DefaultQuery("nickname", "unknown user")

	store.ServeWs(roomid, nickName, hub, ctx.Writer, ctx.Request)
}
