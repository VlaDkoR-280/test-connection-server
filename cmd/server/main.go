package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"test-connection-server/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := server.New()
	handler := server.LoggingMiddleware(app.Mux())

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("🚀 Сервер запущен на http://localhost:%s", port)
	log.Printf("📊 Главная страница: http://localhost:%s/", port)

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("❌ Ошибка запуска сервера: ", err)
	}
}
