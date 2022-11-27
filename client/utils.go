package client

import (
	"log"
	"net/url"

	"github.com/LeoReeYang/im2/models"
	"github.com/gorilla/websocket"
)

func newWebConn(url string) *websocket.Conn {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	return c
}

func newWsUrl(name string, host string, path string) *url.URL {
	u := &url.URL{
		Scheme:   Scheme,
		Host:     host,
		Path:     path,
		RawQuery: "nickname=" + name,
	}
	return u
}

func (c *Client) ListenMsg() {
	for {
		message := models.Message{}
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Println("Message read error :", err)
		}

		c.readBuf <- &message

		// log.Printf("client <%s> recive msg: %v\n", c.Name, message)
	}
}
