package client

import (
	"flag"
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	// var id int

	var cmd string
	// flag.IntVar(&id, "i", 0, "-i input id: ")
	flag.StringVar(&cmd, "c", "123", "-c xxx")

	flag.Parse()

	// args := flag.Args()
	// for _, arg := range args {
	// 	log.Println("arg get :", arg)
	// }

	// fmt.Println("input id : ", id)
	fmt.Println(cmd)
}
