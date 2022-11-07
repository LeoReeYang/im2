package utils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
	Ctx context.Context
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error of InitConfig() : %v\n", err)
	}
}

func InitRedis() {
	Ctx = context.Background()
	addr := viper.GetString("redis.addr")
	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	defer func(*redis.Client) {
		Rdb.Close()
	}(Rdb)
}

func InitDB() {
	dsn := viper.GetString("mysql.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("MySQL open error : %v\n", err)
	}

	DB = db
}
