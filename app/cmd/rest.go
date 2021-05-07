package cmd

import (
	"log"
	"net/http"

	rest "github.com/julianjca/julian-golang-training-beginner/internal/rest"
	httprouter "github.com/julienschmidt/httprouter"
)

func Execute() {
	port := ":3000"

	r := httprouter.New()

	rest.InitHandler(r)
	rest.InitPaymentCodeRESTHandler(r, paymentCodeService)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
