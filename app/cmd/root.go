package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"

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
	paymentCodeRepository *postgres.PaymentCodeRepository
	paymentCodeService    *paymentcode.PaymentCodeService
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Application",
	}
)

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	paymentCodeService = paymentcode.NewService(paymentCodeRepository)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// Execute will call the root command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}