// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/LeoReeYang/im2/models"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func t() {
	var name string
	flag.StringVar(&name, "n", "Guest", "-n xxx")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/testroom", RawQuery: "nickyname=" + name}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("client %s recv msg: %s\n", name, string(message))
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	msg := models.Message{
		Sender:    name,
		Recipient: "server",
		Type:      "message",
		Content:   "test",
		Timestamp: time.Now().Unix(),
	}

	for {
		select {
		case <-done:
			return
		case t, ok := <-ticker.C:
			if !ok {
				log.Fatal("Ticker failed.")
				return
			}

			msg.Timestamp = t.Unix()

			b, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s Sent msg : %s\n", name, string(b))
			c.WriteMessage(websocket.BinaryMessage, b)
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
