package connection

import (
	"fmt"
	"time"

	"github.com/LeoReeYang/im2/models"
	"github.com/LeoReeYang/im2/mq"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

type Status bool

const (
	Online  Status = true
	Offline Status = false
)

type Connection struct {
	name   string
	conn   *websocket.Conn
	alive  Status
	send   chan *models.Message
	recive chan *models.Message
}

func NewConnection(Name string, conn *websocket.Conn) *Connection {
	return &Connection{
		name:   Name,
		conn:   conn,
		alive:  Online,
		send:   make(chan *models.Message, 1024),
		recive: make(chan *models.Message, 1024),
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

func (c *Connection) PutReciveChannel(msg *models.Message) {
	c.recive <- msg
}

func (c *Connection) PutSendChannel(msg *models.Message) {
	c.send <- msg
}

func (c *Connection) GetName() string {
	return c.name
}

func (c *Connection) GetAlive() Status {
	return c.alive
}

func (c *Connection) CloseChannel() {
	c.conn.Close()
}

func (c *Connection) Offline() {
	c.alive = Offline
}

func (c *Connection) read() {
	defer func() {
		c.Offline()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg models.Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			color.Red("Huber Connection ReadJSON:", err)
			return
		}
		mq.MQ <- &msg

		color.Magenta("Client %v get msg from conn:\n msg: %v\n", c.name, msg)
	}
}

func (c *Connection) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Offline()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} else {
				err := c.conn.WriteJSON(message)
				if err != nil {
					fmt.Println("Client WriteJson Error: ", err)
					return
				}

				color.Magenta("Client: < %s > send msg to conn\nmsg : %v\n", c.name, message)
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Connection) ListenChannel() {
	go c.read()
	go c.write()
}
