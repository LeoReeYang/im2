package server

import "im/store"

type Server struct {
	hub *store.Hub

	Messager store.MessageStore
}

func NewServer(ms store.MessageStore) *Server {
	s := new(Server)
	s.hub = store.NewHub()
	s.Messager = ms
	return s
}

func (s *Server) Setup() {
	go s.hub.Run()
}

func (s *Server) GetHub() *store.Hub {
	return s.hub
}
