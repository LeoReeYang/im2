package v1

import (
	"net/http"

	"github.com/LeoReeYang/im2/global"
	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/utils/token"
	"github.com/gin-gonic/gin"
)

func Profile(ctx *gin.Context) {
	user_id, _, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.Get(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Me(ctx *gin.Context) {
	_, name, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hub := global.Server.GetHub()
	conn := hub.UserMeta.Get(name)
	status := conn.GetAlive()
	ctx.JSON(http.StatusOK, gin.H{
		"name":   name,
		"status": status,
	})
}

func All(ctx *gin.Context) {
	hub := global.Server.GetHub()
	users := hub.UserMeta.All()
	ctx.JSON(http.StatusOK, gin.H{
		"online_users": users,
	})
}
