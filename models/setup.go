package models

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := viper.GetString("mysql.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("MySQL initialize error : ", err)
	}

	DB = db

	DB.AutoMigrate(&User{}, &Friend{}, &Block{}, &Message{})
}

func GetDB() *gorm.DB {
	return DB
}
