package v1

import (
	"net/http"

	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/utils/token"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

// QueryUser     godoc
// @Summary      Query a user
// @Description  Responds with the specific user.
// @Tags         books
// @Param        id path string true "The ID"
// @Produce      json
// @Success      200  string  test
// @Router       /user [get]
func Query(ctx *gin.Context) {
	db := models.GetDB()
	user := models.User{}
	user_id, _, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
		return
	}
	// Preload needs the Name of struct field , Not the json name
	db.Preload("Friends").Preload("Blocks").First(&user, user_id)

	if user.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "User Not Found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func Delete(ctx *gin.Context) {
	db := models.GetDB()
	user_id, _, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
		return
	}
	db.Delete(&models.User{}, user_id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Update(ctx *gin.Context) {
	req := models.UpdateNameRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	db := models.GetDB()
	user_id, _, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
		return
	}
	user := models.User{}
	db.Model(&models.User{}).Find(&user, user_id)
	db.Model(&user).Update("name", req.Name)

	ctx.JSON(http.StatusOK, user)
}

func AddFriend(ctx *gin.Context) {
	req := models.AddFriendRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	db := models.GetDB()
	user_id, _, err := token.ExtractTokenInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
		return
	}

	friend_relation := models.Friend{
		UserID:   user_id,
		FriendID: req.FriendID,
	}

	db.Model(&models.Friend{}).Create(&friend_relation)
	if friend_relation.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, "Internal Database error.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "add friend Successfully.",
	})

}
