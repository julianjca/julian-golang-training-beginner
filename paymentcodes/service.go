package paymentcodes

import (
	"fmt"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

type GetByIDResponse struct {
	ID string
}
type repository interface {
	Create(p *golangtraining.PaymentCode) (*golangtraining.PaymentCode, error)
	GetByID(ID string) (golangtraining.PaymentCode, error)
}

type PaymentCodeService struct {
	repo repository
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo repository,
) *PaymentCodeService {
	return &PaymentCodeService{
		repo: repo,
	}
}

func (s PaymentCodeService) Create(p *golangtraining.PaymentCode) error {
	_, err := s.repo.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s PaymentCodeService) GetByID(ID string) (res golangtraining.PaymentCode, err error) {
	res, err = s.repo.GetByID(ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
