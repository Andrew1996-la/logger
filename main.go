package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/error", errorHandler)

	logger.Info("Server started", zap.String(
		"port", "8080",
	))

	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
	logger.Info("hello handler",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
	logger.Error("error handler",
		zap.String("method", r.Method),
	)
}
