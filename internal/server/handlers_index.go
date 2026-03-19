package server

import (
	_ "embed"
	"net/http"
)

//go:embed static/index.html
var indexHTML []byte

func (s *Server) indexHandler(w http.ResponseWriter, _ *http.Request) {
	s.IncrementRequestCount()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(indexHTML)
}
