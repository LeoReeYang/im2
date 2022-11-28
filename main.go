package main

import (
	"github.com/LeoReeYang/im2/global"
	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/router"
	"github.com/LeoReeYang/im2/utils"
)

func main() {
	utils.InitConfig()
	models.InitDB()
	global.Init()

	router.SetupRouters()
}
