package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/error", errorHandler)

	fmt.Println("Server start on 8080 port")
	http.ListenAndServe(":8080", mux)
}
