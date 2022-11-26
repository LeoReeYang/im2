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
	Conn    *websocket.Conn
	readBuf chan *models.Message
}

func NewClient(name string, host string, path string) *Client {
	u := newWsUrl(name, host, path)
	client := &Client{
		Name:    name,
		Conn:    newWebConn(u.String()),
		readBuf: make(chan *models.Message, 1024),
	}

	// listen Message from the conn
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

	err := c.Conn.WriteJSON(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("<%s> Sent msg : %v\n", c.Name, msg)
}

func (c *Client) Receive() (*models.Message, bool) {
	// select {
	// case msg := <-c.readBuf:
	// 	return msg, true
	// default:
	// 	return nil, false
	// }
	defer func() (*models.Message, bool) {
		return nil, false
	}()

	for {
		msg := <-c.readBuf
		return msg, true
	}
}

func (c *Client) AddFriend() {}

func (c *Client) BlockFriend() {}
