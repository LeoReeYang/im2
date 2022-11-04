package mysql

import (
	"log"

	"github.com/spf13/viper"
)

func T() {
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file :%v", err)
	}

}
