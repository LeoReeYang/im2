package server

import (
	"github.com/LeoReeYang/im2/store"
)

type Server struct {
	// hub   *store.Hub
	huber *store.Huber

	Messager store.MessageStore
	// router   *gin.Engine
}

func NewServer(ms store.MessageStore) *Server {
	s := new(Server)
	// s.hub = store.NewHub()
	s.huber = store.NewHuber()
	s.Messager = ms
	return s
}

func (s *Server) Setup() {
	// go s.hub.Run()
	go s.huber.Run()
}

func (s *Server) GetHub() *store.Huber {
	return s.huber
}
