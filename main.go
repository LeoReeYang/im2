package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"im/models"
	"im/server"
	"im/store"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// err = c.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
		msg := models.Message{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("echo msg : %v\n", msg)
	}
}

func main() {
	// router.SetupRouters()
	flag.Parse()

	r := gin.Default()

	server := server.NewServer(store.NewMsgQueue())

	hub := server.GetHub()
	go hub.Run()

	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})

	r.GET("/Chat", func(ctx *gin.Context) {
		roomid := ctx.DefaultQuery("roomid", "001")
		nickName := ctx.DefaultQuery("nickname", "unknown user")

		// roomid := "001"
		// nickName := "unknown user"

		store.ServeWs(roomid, nickName, hub, ctx.Writer, ctx.Request)
	})

	r.Run(*addr)
}
