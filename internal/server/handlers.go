package server

import (
	"net/http"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check\n"))
}

func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) SecondPage(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) ThirdPage(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) FourthPage(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
