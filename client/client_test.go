package client

import (
	"testing"
	"time"

	"github.com/fatih/color"
)

func TestConn(t *testing.T) {
	url := newWsUrl("cy", *Addr, Path)

	c := newWebConn(url.String())

	if c == nil {
		t.Errorf("conn is nill")
	}
}

func TestSendAndRecive(t *testing.T) {
	Me := NewClient("Me", *Addr, Path)

	want := string("hi!")

	Me.Send("Cy", want)

	if msg, ok := Me.Receive(); ok {
		color.Cyan("Message get:", msg.Content)
		if msg.Content != want {
			t.Errorf("content not match.")
		}
	}
}

func Test2People(t *testing.T) {
	Alice := NewClient("Alice", *Addr, Path)
	Bob := NewClient("Bob", *Addr, Path)

	want := string("hello, Bob!")
	Alice.Send("Bob", want)

	if msg, ok := Bob.Receive(); ok {
		color.Yellow("Bob gets message:", msg.Content)
		if msg.Content != want {
			t.Errorf("content doesn't match.")
		}
	}

	time.Sleep(time.Second)

	Bob.Send("Alice", "hi, Alice!")

	if msg, ok := Alice.Receive(); ok {
		color.Blue("Alice gets message:", msg.Content)
		if msg.Content != "hi, Alice!" {
			t.Errorf("content not match.")
		}
	}
}
