package server

import (
	"github.com/LeoReeYang/im2/store"
)

type Server struct {
	huber *store.Huber
}

func NewServer(ms store.MessageStore) *Server {
	s := new(Server)
	s.huber = store.NewHuber(ms)
	return s
}

func (s *Server) Setup() {
	go s.huber.Run()
}

func (s *Server) GetHub() *store.Huber {
	return s.huber
}
