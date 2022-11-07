package client

import (
	"flag"
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	var id int

	flag.IntVar(&id, "i", 0, "-i input id: ")

	flag.Parse()

	fmt.Println("input id : ", id)
}
