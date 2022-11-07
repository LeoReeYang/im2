package v1

import (
	"net/http"

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
func QueryUser(ctx *gin.Context) {
	// db := mysql.GetDB()
	ctx.JSON(http.StatusOK, "helloworld")
}

func DeleteUser(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {}
