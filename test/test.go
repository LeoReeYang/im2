package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string
	// flag.IntVar(&id, "i", 0, "-i input id: ")
	flag.StringVar(&cmd, "c", "123", "-c xxx")

	flag.Parse()

	// args := flag.Args()
	// for _, arg := range args {
	// 	fmt.Println("arg get :", arg)
	// }
	fmt.Println(cmd)
}
