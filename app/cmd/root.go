package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	postgres "github.com/julianjca/julian-golang-training-beginner/internal/postgres"
	starwars "github.com/julianjca/julian-golang-training-beginner/internal/starwars"
	paymentcode "github.com/julianjca/julian-golang-training-beginner/paymentcodes"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "golang_training"
)

var (
	paymentCodeRepository *postgres.PaymentCodeRepository
	paymentCodeService    *paymentcode.PaymentCodeService
)

func init() {
	httpClient := initHttpClient()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	starwarsClient := starwars.NewStarWarsClient(httpClient)
	paymentCodeService = paymentcode.NewService(paymentCodeRepository, *starwarsClient)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

func initHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}
