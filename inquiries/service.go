package paymentcodes

import (
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	paymentcode "github.com/julianjca/julian-golang-training-beginner/paymentcodes"
)

//go:generate mockgen -destination=mocks/mock_paymentcodes_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error)
}

type InquiryService struct {
	repo               Repository
	paymentCodeService paymentcode.PaymentCodeService
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
	paymentCodeService paymentcode.PaymentCodeService,
) *InquiryService {
	return &InquiryService{
		repo:               repo,
		paymentCodeService: paymentCodeService,
	}
}

func (i InquiryService) Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error) {
	// check if payment code exist
	_, err := i.paymentCodeService.GetByPaymentCode(p.PaymentCode)
	if err != nil {
		return nil, err
	}

	res, err := i.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
