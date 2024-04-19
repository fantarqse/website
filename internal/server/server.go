package server

import (
	"html/template"
	"net/http"
)

type Server struct {
	mux       *http.ServeMux
	templates *template.Template
	// TODO: Add configuration
	// TODO: Add logger
}

func New(tmpls *template.Template) *Server {
	return &Server{
		mux:       http.NewServeMux(),
		templates: tmpls,
	}
}

func (s *Server) Routes() {
	// static files
	/*
		TODO: read about http.FileServer
		TODO: read about http.StripPrefix
	*/
	fs := http.FileServer(http.Dir("./web/static"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", fs))

	s.mux.HandleFunc("GET /index", s.MainPage)
	s.mux.HandleFunc("GET /health", s.Health)
}

func (s *Server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}
