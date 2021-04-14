package main

import (
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"healthy"}`))

	w.WriteHeader(http.StatusOK)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"hello world"}`))

	w.WriteHeader(http.StatusOK)
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
