package client

import (
	"fmt"
	"log"
	"time"

	"github.com/LeoReeYang/im2/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Name    string
	C       *websocket.Conn
	readBuf chan *models.Message
}

func (c *Client) ListenMsg() {
	for {
		message := models.Message{}
		err := c.C.ReadJSON(&message)
		if err != nil {
			// log.Fatal("Message read error :", err)
			log.Println("Message read error :", err)
		}

		c.readBuf <- &message

		log.Printf("client <%s> recive msg: %v\n", c.Name, message)
	}
}

func NewClient(name string, host string, path string) *Client {
	u := newWsUrl(name, host, path)
	client := &Client{
		Name:    name,
		C:       newWebConn(u.String()),
		readBuf: make(chan *models.Message, 1024),
	}

	go client.ListenMsg()

	return client
}

func (c *Client) Send(recipient, content string) {
	msg := &models.Message{
		Sender:    c.Name,
		Recipient: recipient,
		Type:      "message",
		Content:   content,
		Timestamp: time.Now().Unix(),
	}

	err := c.C.WriteJSON(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("<%s> Sent msg : %v\n", c.Name, msg)
}

func (c *Client) Receive() (*models.Message, bool) {
	select {
	case msg := <-c.readBuf:
		return msg, true
	default:
		return nil, false
	}
}

func (c *Client) AddFriend() {}

func (c *Client) BlockFriend() {}
