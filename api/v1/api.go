package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/store"
	"github.com/fatih/color"
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
	ctx.JSON(http.StatusOK, "helloworld")
}

func ServerWs(ctx *gin.Context, hub *store.Hub) {
	roomid := ctx.DefaultQuery("roomid", "001")
	nickName := ctx.DefaultQuery("nickname", "unknown user")

	store.ServeWs(roomid, nickName, hub, ctx.Writer, ctx.Request)
}

func Test(ctx *gin.Context) {

	// room := ctx.Param("room")
	name := ctx.Query("nickyname")
	// user := ctx.PostForm("user")

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// client := store.NewClient(user, "test", name, conn, hub)

	go func() {
		for {
			var msg models.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			color.Cyan("ussr :%v", name)
			color.Cyan("Message :%v", msg)

			b, _ := json.Marshal(msg)
			err = conn.WriteMessage(websocket.BinaryMessage, b)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
