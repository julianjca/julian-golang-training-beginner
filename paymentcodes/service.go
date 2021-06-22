package paymentcodes

import (
	"fmt"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

type GetByIDResponse struct {
	ID string
}

//go:generate mockgen -destination=mocks/mock_paymentcodes_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.PaymentCode) (*golangtraining.PaymentCode, error)
	GetByID(ID string) (golangtraining.PaymentCode, error)
}

//go:generate mockgen -destination=mocks/mock_starwars_client.go -package=mocks . StarWarsClient
type StarWarsClient interface {
	GetCharacters() (*golangtraining.StarWarsResponse, error)
}

type PaymentCodeService struct {
	repo           Repository
	starwarsClient StarWarsClient
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
	starwarsClient StarWarsClient,
) *PaymentCodeService {
	return &PaymentCodeService{
		repo:           repo,
		starwarsClient: starwarsClient,
	}
}

func (s PaymentCodeService) Create(p *golangtraining.PaymentCode) error {
	_, err := s.starwarsClient.GetCharacters()
	_, err = s.repo.Create(p)
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
