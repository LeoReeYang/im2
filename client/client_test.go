package client

import (
	"fmt"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	url := newWsUrl("cy", *Addr, Path)

	c := newWebConn(url.String())

	if c == nil {
		t.Errorf("conn is nill")
	}
}

func TestRecv(t *testing.T) {
	Client := NewClient("yzy", *Addr, Path)
	// time.Sleep(time.Second)

	if msg, ok := Client.Receive(); ok {
		fmt.Println("Message get:", msg)
		// if msg.Content != "民族富强！" {
		// 	t.Errorf("content not match.")
		// }
	}
	time.Sleep(10 * time.Second)
}

func TestSendMessage(t *testing.T) {
	Client := NewClient("test", *Addr, Path)
	time.Sleep(time.Second)

	Client.Send("yzy", "Hello!")
	time.Sleep(5 * time.Second)
}

func TestSendAndRecive(t *testing.T) {
	ClientA := NewClient("cy", *Addr, Path)

	want := string("民族富强！")

	ClientA.Send("cy", want)
	time.Sleep(time.Second)

	if msg, ok := ClientA.Receive(); ok {
		fmt.Println("Message get:", msg.Content)
		if msg.Content != want {
			t.Errorf("content not match.")
		}
	}
}

func TestClient(t *testing.T) {
	client := NewClient("Alice", *Addr, Path)
	client.ListenMsg()

	defer client.C.Close()

	want := "test mgs"

	client.Send("server", want)

	msg, ok := client.Receive()

	if ok {
		fmt.Println("msg get:", msg)
	}

	if msg.Content != "test msg" {
		t.Errorf("get msg: %v, but want: %s", msg.Content, want)
	}
}

func Test2People(t *testing.T) {
	clientB := NewClient("Bob", *Addr, Path)
	defer time.Sleep(time.Second)

	clientA := NewClient("Alice", *Addr, Path)

	// defer clientB.C.Close()
	// defer clientA.C.Close()

	want := string("hello!")
	clientA.Send("Bob", want)

	time.Sleep(time.Second)

	if msg, ok := clientB.Receive(); ok {
		fmt.Println("Message get:", msg.Content)
		if msg.Content != want {
			t.Errorf("content not match.")
		}
	}

	clientB.Send("Alice", "hi!")
	time.Sleep(time.Second)

	if msg, ok := clientA.Receive(); ok {
		fmt.Println("Message get:", msg.Content)
		if msg.Content != "hi!" {
			t.Errorf("content not match.")
		}
	}
}

func TestMessageSend(t *testing.T) {

}
