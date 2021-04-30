package cmd

import (
	"log"
	"net/http"

	rest "github.com/julianjca/julian-golang-training-beginner/internal/rest"
)

func Execute() {
	port := ":3000"

	rest.InitHandler()

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
