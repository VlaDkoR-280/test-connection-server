package server

import (
	"encoding/json"
	"net/http"
	"test-connection-server/internal/health"
)

func (s *Server) apiHealthHandler(w http.ResponseWriter, _ *http.Request) {
	s.IncrementRequestCount()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resp := health.BuildResponse(s.Uptime(), s.RequestCount())
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) pingHandler(w http.ResponseWriter, _ *http.Request) {
	s.IncrementRequestCount()
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("pong"))
}
