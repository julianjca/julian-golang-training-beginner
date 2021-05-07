package golangtraining

import (
	"time"
)

type PaymentCode struct {
	ID             string    `json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name" validate:"required"`
	Status         string    `json:"status"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type IPaymentCodeService interface {
	Create(p *PaymentCode) error
	GetByID(ID string) (res PaymentCode, err error)
}

type IPaymentCodeRepository interface {
	Create(p *PaymentCode) (res *PaymentCode, err error)
	GetByID(ID string) (res PaymentCode, err error)
}
