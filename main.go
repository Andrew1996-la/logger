package main

import (
	"net/http"

	"go.uber.org/zap"
	"logger/logger"
)


func main() {
	logger.Init(logger.Config{
		Env: "dev",
	})
	
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
		logger.String("method", r.Method),
		logger.String("path", r.URL.Path),
	)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
	logger.Error("error handler",
		logger.String("method", r.Method),
	)
}
