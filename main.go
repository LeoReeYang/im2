package main

import (
	"github.com/LeoReeYang/im2/mysql"
	"github.com/LeoReeYang/im2/router"
)

func main() {
	mysql.InitDB()
	router.SetupRouters()
}
