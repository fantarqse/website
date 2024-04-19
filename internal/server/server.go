package server

import "net/http"

type Server struct {
	mux *http.ServeMux
	// TODO: Add configuration
	// TODO: Add logger
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /health", s.Health)
}

func (s *Server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}
