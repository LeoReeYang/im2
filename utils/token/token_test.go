package token

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func TestSecret(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error of InitConfig() : %v\n", err)
	}

	Secret := viper.GetInt32("Secret.life_span")

	log.Println(Secret)
}

func TestTokenSign(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error of InitConfig() : %v\n", err)
	}

	secret := []byte(viper.GetString("secret.key"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":       "bar",
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"Issuer":    "test",
		"Subject":   "somebody",
		"ID":        "1",
		"Audience":  []string{"somebody_else"},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		t.Fatal(err)
	}

	// log.Println(tokenString, err)

	new_token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := new_token.Claims.(jwt.MapClaims); ok && new_token.Valid {
		fmt.Println(claims["foo"], claims["ExpiresAt"])
	}
}
