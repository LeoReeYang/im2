package models

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAddFriend(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error of InitConfig() : %v\n", err)
	}
	dsn := viper.GetString("mysql.dsn")

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("MySQL initialize error : ", err)
	}

	cy := User{
		Name:     "cy",
		Password: "123",
		Friends: []Friend{
			{UserID: 1, FriendID: 2},
		},
	}
	me := User{
		Name:     "yzy",
		Password: "123",
		Friends: []Friend{
			{UserID: 2, FriendID: 1},
		},
	}
	DB.Create(&cy)
	DB.Create(&me)

	if cy.ID == 0 && me.ID == 0 {
		t.Errorf("ID = 0")
	}
}
