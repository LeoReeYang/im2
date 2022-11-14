package client

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func newWebConn(url string) *websocket.Conn {
	// log.Println(url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		// log.Fatal("Dial :", err)
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
		RawQuery: "nickyname=" + name,
	}
	return u
}
