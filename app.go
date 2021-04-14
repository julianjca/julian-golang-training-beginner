package main

import (
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"status":"healthy"}`)
	w.Write(jsonData)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"message":"hello world"}`)
	w.Write(jsonData)
}

func handleRequests() {
	port := ":3000"
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/hello-world", helloWorld)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	handleRequests()
}
