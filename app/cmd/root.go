package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/julianjca/julian-golang-training-beginner/internal/sqs"

	"github.com/julianjca/julian-golang-training-beginner/app/cmd/helpers"

	"github.com/julianjca/julian-golang-training-beginner/internal/jobs"
	"github.com/spf13/cobra"

	"github.com/julianjca/julian-golang-training-beginner/inquiries"
	"github.com/julianjca/julian-golang-training-beginner/internal/postgres"
	"github.com/julianjca/julian-golang-training-beginner/paymentcodes"
	"github.com/julianjca/julian-golang-training-beginner/payments"

	_ "github.com/lib/pq"
)

var (
	paymentCodeRepository *postgres.PaymentCodeRepository
	paymentCodeService    *paymentcodes.PaymentCodeService
	inquiriesRepository   *postgres.InquiriesRepository
	inquiriesService      *inquiries.InquiryService
	paymentsService       *payments.PaymentService
	paymentsRepository    *postgres.PaymentsRepository
	rootCmd               = &cobra.Command{
		Use:   "app",
		Short: "Application",
	}
	expirePaymentCodesJob jobs.ExpirePaymentCodesJob
)

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	host := helpers.MustHaveEnv("POSTGRES_HOST")
	portStr := helpers.MustHaveEnv("POSTGRES_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err, "POSTGRES_PORT is not well set ")
	}
	user := helpers.MustHaveEnv("POSTGRES_USER")
	password := helpers.MustHaveEnv("POSTGRES_PASSWORD")
	dbname := helpers.MustHaveEnv("POSTGRES_DB_NAME")

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqsPublisher := initSQSPublisher()

	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	paymentCodeService = paymentcodes.NewService(paymentCodeRepository)
	inquiriesRepository = postgres.NewInquiriesRepository(db)
	inquiriesService = inquiries.NewService(inquiriesRepository, *paymentCodeService)
	paymentsRepository = postgres.NewPaymentsRepository(db)
	paymentsService = payments.NewService(paymentsRepository, *inquiriesService, sqsPublisher)

	fmt.Println("Successfully connected!")

	expirePaymentCodesJob = jobs.ExpirePaymentCodesJob{PaymentCodesService: paymentCodeService}
}

// Execute will call the root command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initSQSPublisher() *sqs.Publisher {
	region := helpers.MustHaveEnv("SQS_AWS_REGION")
	endpoint := helpers.MustHaveEnv("SQS_ENDPOINT")

	s, err := session.NewSession(&aws.Config{
		Region:   &region,
		Endpoint: &endpoint,
	})
	if err != nil {
		panic(err)
	}

	q := helpers.MustHaveEnv("SQS_QUEUE_NAME")
	p, err := sqs.NewPublisher(s, q)
	if err != nil {
		panic(err)
	}

	return p
}
