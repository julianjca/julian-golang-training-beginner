package golangtraining

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type PaymentCode struct {
	ID             string    `json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExpirationDate string    `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type IPaymentCodeService interface {
	Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type IPaymentCodeRepository interface {
	Create(p *PaymentCode) (res PaymentCode, err error)
	GetByID(ID string) (res PaymentCode, err error)
}
