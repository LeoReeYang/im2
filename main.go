package main

import (
	"github.com/LeoReeYang/im2/router"
	"github.com/LeoReeYang/im2/utils"
)

func main() {
	utils.InitConfig()
	utils.InitRedis()
	utils.InitDB()
	router.SetupRouters()
}
