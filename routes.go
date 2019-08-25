package main

func (s *Server) routes() {
	s.Router.HandleFunc("/", s.handleIndex())
	s.Router.HandleFunc("/all", s.handleAll(s.db))
}
