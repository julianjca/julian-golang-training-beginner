package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	rest "github.com/julianjca/julian-golang-training-beginner/internal/rest"
	"github.com/julianjca/julian-golang-training-beginner/internal/starwars"
	"github.com/julienschmidt/httprouter"
)

var restCommand = &cobra.Command{
	Use:   "rest",
	Short: "Start REST server",
	Run:   restServer,
}

func init() {
	rootCmd.AddCommand(restCommand)
}

func restServer(cmd *cobra.Command, args []string) {
	port := ":3000"

	r := httprouter.New()

	httpClient := initHttpClient()
	rest.InitHandler(r)
	starwarsClient := starwars.NewStarWarsClient(httpClient, "https://swapi.dev/api/people")

	rest.InitPaymentCodeRESTHandler(r, paymentCodeService, starwarsClient)

	rest.InitInquiryRESTHandler(r, inquiriesService)

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
