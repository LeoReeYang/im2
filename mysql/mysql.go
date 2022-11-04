package mysql

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error of config file :%v\n", err)
	}
	dsn := viper.GetString("mysql.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("MySQL open error : %v\n", err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
