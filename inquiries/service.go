package paymentcodes

import (
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

//go:generate mockgen -destination=mocks/mock_paymentcodes_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error)
}

type InquiryService struct {
	repo Repository
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
) *InquiryService {
	return &InquiryService{
		repo: repo,
	}
}

func (i InquiryService) Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error) {
	res, err := i.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
