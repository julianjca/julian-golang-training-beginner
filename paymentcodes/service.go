package paymentcodes

import (
	"encoding/json"
	"fmt"
	"net/http"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

type CreateResponse struct {
	Status string
}

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

func (s service) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	paymentCodeBody := &golangtraining.PaymentCode{
		Status: "ACTIVE",
	}

	w.WriteHeader(http.StatusOK)
	s.repo.Create(paymentCodeBody)
	jsonData := CreateResponse{Status: "123"}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)

}

func (s service) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	s.repo.GetByID("123")
	jsonData := GetByIDResponse{ID: ps.ByName("id")}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}
