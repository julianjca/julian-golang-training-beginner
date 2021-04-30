package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := &HealthResponse{Status: "healthy"}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := &HelloResponse{Message: "hello world"}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}

func InitHandler() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/hello-world", helloWorld)
}
