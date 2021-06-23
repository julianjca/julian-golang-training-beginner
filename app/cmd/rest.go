package cmd

import (
	"log"
	"net/http"
	"time"

	rest "github.com/julianjca/julian-golang-training-beginner/internal/rest"
	"github.com/julianjca/julian-golang-training-beginner/internal/starwars"
	"github.com/julienschmidt/httprouter"
)

func Execute() {
	port := ":3000"

	r := httprouter.New()

	httpClient := initHttpClient()
	rest.InitHandler(r)
	starwarsClient := starwars.NewStarWarsClient(httpClient)

	rest.InitPaymentCodeRESTHandler(r, paymentCodeService, starwarsClient)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func initHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}
