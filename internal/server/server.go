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

func (s *Server) Routes() *Server {
	// static files
	/*
		TODO: read about http.FileServer
		TODO: read about http.StripPrefix
	*/
	fs := http.FileServer(http.Dir("./web/static"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", fs))

	s.mux.HandleFunc("GET /index", s.MainPage)    // TODO: change the handler name
	s.mux.HandleFunc("GET /second", s.SecondPage) // TODO: change the handler name
	s.mux.HandleFunc("GET /third", s.ThirdPage)   // TODO: change the handler name
	s.mux.HandleFunc("GET /fourth", s.FourthPage) // TODO: change the handler name

	s.mux.HandleFunc("GET /health", s.Health)

	return s
}

func (s *Server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}
