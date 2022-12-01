package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GenerateToken(user_id uint, user_name string) (string, error) {
	token_lifespan := viper.GetInt32("secret.life_span")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["username"] = user_name
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("secret.key")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("secret.key")), nil
	})

	return token, err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return bearerToken
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := ParseToken(tokenString)
	return err
}

func ExtractTokenInfo(c *gin.Context) (uint, string, error) {
	tokenString := ExtractToken(c)

	token, err := ParseToken(tokenString)
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		user_name := fmt.Sprintf("%v", claims["user_name"])
		if err != nil {
			return 0, "", err
		}
		return uint(uid), user_name, nil
	}
	return 0, "", nil
}
