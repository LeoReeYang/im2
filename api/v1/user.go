package v1

import (
	"net/http"

	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/utils/token"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

// var DB *gorm.DB

func CreateUser(ctx *gin.Context) {
	// db := mysql.GetDB()
	// db.Create(user)
}

// GetUser       godoc
// @Summary      Get a user
// @Description  Responds with the specific user.
// @Tags         books
// @Param        id path string true "The ID"
// @Produce      json
// @Success      200  string  test
// @Router       /user [get]
func QueryUser(c *gin.Context) {
	// db := mysql.GetDB()
	// c.JSON(http.StatusOK, "helloworld")

	// data := LoginUser{}
	// err := c.ShouldBindJSON(&data)
	// if err != nil {
	// 	c.JSON(http.StatusNoContent, err)
	// }
}

func DeleteUser(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {}

// func

func CurrentUser(ctx *gin.Context) {
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
