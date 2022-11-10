package client

import (
	"encoding/json"
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
	// go func() {
	for {
		_, message, err := c.C.ReadMessage()
		if err != nil {
			log.Fatal("Message read error :", err)
		}

		msg := models.NewMessage()
		json.Unmarshal(message, msg)
		c.readBuf <- msg

		// log.Printf("client <%s> recive msg: %s\n", c.Name, string(message))
	}
	// }()
}

func NewClient(name string, host string, path string) *Client {
	u := newWsUrl(name, host, path)
	// c := newWebConn(u.String())
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
		ID:        "testroom",
		Sender:    c.Name,
		Recipient: recipient,
		Type:      "message",
		Content:   content,
		Timestamp: time.Now().Unix(),
	}

	b, err := json.Marshal(*msg)
	if err != nil {
		log.Println("Message json.Marshal err: ", err)
	}
	fmt.Printf("<%s> Sent msg : %s\n", c.Name, b)
	c.C.WriteMessage(websocket.BinaryMessage, b)
}

func (c *Client) Receive() (*models.Message, bool) {
	select {
	case msg := <-c.readBuf:
		// fmt.Println("Message Get:", msg)
		return msg, true
	default:
		return nil, false
	}
}

func (c *Client) AddFriend() {}

func (c *Client) BlockFriend() {}
