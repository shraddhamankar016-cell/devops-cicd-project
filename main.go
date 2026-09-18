package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from my DevOps CI/CD project! Version 1.0")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
