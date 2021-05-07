package paymentcodes

import (
	"fmt"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

type GetByIDResponse struct {
	ID string
}

type service struct {
	repo golangtraining.IPaymentCodeRepository
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo golangtraining.IPaymentCodeRepository,
) golangtraining.IPaymentCodeService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(p *golangtraining.PaymentCode) error {
	_, err := s.repo.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetByID(ID string) (res golangtraining.PaymentCode, err error) {
	res, err = s.repo.GetByID(ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
