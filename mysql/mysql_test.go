package mysql

import (
	"log"
	"testing"

	"github.com/LeoReeYang/im2/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestMysql(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file :%v", err)
	}
	dsn := viper.GetString("mysql.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql error %v", err)
	}

	db.AutoMigrate(&models.User{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1) // 根据整型主键查找

	log.Println(product.Price)
}
