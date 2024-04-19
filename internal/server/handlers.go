package server

import "net/http"

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check\n"))
}
