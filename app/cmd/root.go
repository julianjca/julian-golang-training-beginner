package cmd

import (
	"database/sql"
	"fmt"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	postgres "github.com/julianjca/julian-golang-training-beginner/internal/postgres"
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
	paymentCodeRepository golangtraining.IPaymentCodeRepository
	paymentCodeService    golangtraining.IPaymentCodeService
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	paymentCodeService = paymentcode.NewService(paymentCodeRepository)

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
