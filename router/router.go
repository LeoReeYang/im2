package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouters() {
	r := gin.Default()

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
