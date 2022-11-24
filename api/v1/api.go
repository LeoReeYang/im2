package v1

import (
	"log"
	"time"

	"github.com/LeoReeYang/im2/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// @BasePath

// PingExample godoc
// @Summary echo example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /test/echo [get]

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Echo(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	parm := ctx.Query("test")
	bi := []byte(parm)

	log.Println(bi)

	go func() {
		for {
			msg := models.Message{}

			err = c.ReadJSON(&msg)
			if err != nil {
				log.Println("read error: ", err)
			}

			msg.Sender = "server"
			msg.Content = "pong"
			msg.Recipient = "postman"
			msg.Timestamp = time.Now().Unix()

			err = c.WriteJSON(msg)
			if err != nil {
				log.Println("write error: ", err)
			}
		}
	}()
}
