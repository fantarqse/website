package server

import (
	"net/http"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check\n"))
}

func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.Execute(w, "index.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
