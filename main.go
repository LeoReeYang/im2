package main

import (

	// "image/color"

	"github.com/LeoReeYang/im2/router"
	"github.com/fatih/color"
)

var Red = color.New(color.FgRed).FprintfFunc()
var Blue = color.New(color.FgBlue).FprintfFunc()

func main() {
	// utils.InitConfig()
	// utils.InitRedis()
	// utils.InitDB()
	router.SetupRouters()

	// r := gin.Default()

	// r.GET("/ws/:room", func(ctx *gin.Context) {
	// 	w := ctx.Writer
	// 	r := ctx.Request
	// 	room := ctx.Param("room")
	// 	name := ctx.Query("nickyname")
	// 	handler(w, r, room, name)
	// })

	// r.Run(":8080")
}
