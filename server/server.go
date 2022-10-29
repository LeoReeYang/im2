package server

import "im/store"

type Server struct {
	hub *store.Hub

	Messager *store.Messager
}

func NewServer() *Server {
	return &Server{
		hub:      store.NewHub(),
		Messager: store.NewMessager(),
	}
}

func (s *Server) Setup() {
	go s.hub.Run()
}

func (s *Server) GetHub() *store.Hub {
	return s.hub
}
