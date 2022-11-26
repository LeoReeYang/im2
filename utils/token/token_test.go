package token

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestToken(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error of InitConfig() : %v\n", err)
	}

	secrat := viper.GetInt32("secrat.life_span")

	log.Println(secrat)
}
