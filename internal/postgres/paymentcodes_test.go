package postgres

import (
	"os"
	"testing"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type paymentCodesTestSuite struct {
	Suite
}

func TestSuitePaymentCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite for Bank Connector Mapping Repository")
	}
	dsn := os.Getenv("POSTGRES_TEST_URL")
	if dsn == "" {
		dsn = DefaultTestDsn
	}
	bankConnectorMappingSuite := &paymentCodesTestSuite{
		Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}
	suite.Run(t, bankConnectorMappingSuite)
}

func (s paymentCodesTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}
func (s paymentCodesTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func CreateMockPaymentCode() (mockVaSettings golangtraining.PaymentCode) {
	return golangtraining.PaymentCode{
		ID:          "1a510335-83eb-49f4-a121-9dc2d7a11348",
		Name:        "Name",
		Status:      "ACTIVE",
		PaymentCode: "payment-code-1",
	}
}

func (s paymentCodesTestSuite) TestCreatePaymentCode() {
	repo := NewPaymentCodeRepository(s.DBConn)
	paymentCode := CreateMockPaymentCode()
	testCases := []struct {
		desc        string
		expectedErr error
		body        *golangtraining.PaymentCode
		expectedRes *golangtraining.PaymentCode
	}{
		{
			desc:        "insert-success",
			expectedErr: nil,
			body:        &paymentCode,
			expectedRes: &paymentCode,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			res, err := repo.Create(tC.body)
			s.Require().Equal(tC.expectedRes.Name, res.Name)
			s.Require().Equal(tC.expectedRes.PaymentCode, res.PaymentCode)
			s.Require().Equal(tC.expectedRes.Status, res.Status)

			s.Require().Equal(tC.expectedErr, errors.Cause(err))
		})
	}
}

func (s paymentCodesTestSuite) TestGetPaymentCodeByID() {
	repo := NewPaymentCodeRepository(s.DBConn)
	subCompanyCodes := CreateMockPaymentCode()
	testCases := []struct {
		desc        string
		expectedErr error
		body        *golangtraining.PaymentCode
	}{
		{
			desc:        "get-success",
			expectedErr: nil,
			body:        &subCompanyCodes,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			p, err := repo.Create(tC.body)
			_, err = repo.GetByID(p.ID)
			s.Require().Equal(tC.expectedErr, errors.Cause(err))
		})
	}
}
