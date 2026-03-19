package server

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware логирует входящие запросы и время обработки.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s от %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("← %s %s - %v", r.Method, r.URL.Path, time.Since(start))
	})
}
