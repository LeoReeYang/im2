package client

import (
	"flag"
)

var (
	Addr   = flag.String("addr", "localhost:8080", "http service address")
	Scheme = "ws"
	Path   = "/ws/chat"
)
