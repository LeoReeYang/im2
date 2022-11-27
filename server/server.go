package server

type Server struct {
	hub *Hub
}

func NewServer(mh MessageHandler, uh ConnectionHandeler) *Server {
	s := new(Server)
	s.hub = NewHub(mh, uh)

	s.Setup()

	return s
}

func (s *Server) Setup() {
	go s.hub.Run()
}

func (s *Server) GetHub() *Hub {
	return s.hub
}
