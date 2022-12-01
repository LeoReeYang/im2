package v1

import (
	"net/http"

	"github.com/LeoReeYang/im2/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.LoginRequest
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.LoginRsp{
			Message: err.Error(),
		})
		return
	}

	u := models.User{
		Name:     data.Name,
		Password: data.Password,
	}
	// verify the user with a token back when user is valid
	token, err := models.UserVerify(u.Name, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var data models.RegisterRequest
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RegisterRsp{
			Message: err.Error(),
		})
		return
	}

	u := models.User{
		Name:     data.Name,
		Password: data.Password,
	}
	_, err = u.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegisterRsp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, models.RegisterRsp{
		Message: "Successful.",
	})
}
