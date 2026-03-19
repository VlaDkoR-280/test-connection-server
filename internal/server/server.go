package server

import (
	"net/http"
	"sync"
	"time"
)

// Server — HTTP-сервер приложения с счётчиком запросов и временем старта.
type Server struct {
	startTime    time.Time
	requestCount int64
	mu           sync.RWMutex
}

// New создаёт новый экземпляр Server.
func New() *Server {
	return &Server{
		startTime: time.Now(),
	}
}

// IncrementRequestCount увеличивает счётчик запросов (потокобезопасно).
func (s *Server) IncrementRequestCount() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requestCount++
}

// RequestCount возвращает текущее значение счётчика запросов.
func (s *Server) RequestCount() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.requestCount
}

// Uptime возвращает время работы сервера.
func (s *Server) Uptime() time.Duration {
	return time.Since(s.startTime)
}

// Mux возвращает настроенный *http.ServeMux с зарегистрированными маршрутами.
func (s *Server) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.indexHandler)
	mux.HandleFunc("/api/health", s.apiHealthHandler)
	mux.HandleFunc("/ping", s.pingHandler)
	return mux
}
