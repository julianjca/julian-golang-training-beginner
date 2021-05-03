package postgres

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/pkg/errors"
)

type paymentCodeRepository struct {
	DB *sql.DB
}

func NewPaymentCodeRepository(db *sql.DB) golangtraining.IPaymentCodeRepository {
	return &paymentCodeRepository{
		DB: db,
	}
}

func (t paymentCodeRepository) Create(p *golangtraining.PaymentCode) (res golangtraining.PaymentCode, err error) {
	newUUID, err := uuid.NewRandom()

	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return
	}

	p.ID = newUUID.String()
	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now
	p.Status = "ACTIVE"
	p.Name = "lechsa"
	p.PaymentCode = "abc123"

	fmt.Println(t.DB)

	// query := sq.
	// 	Insert("payment_code").Columns("id", "payment_code", "name", "status", "expiration_date", "created_at", "updated_at").
	// 	Values(p.ID, p.PaymentCode, p.Name, p.Status, p.ExpirationDate, p.CreatedAt, p.UpdatedAt).PlaceholderFormat(sq.Dollar)

	query := sq.
		Insert("payment_codes").
		Columns("id", "payment_code", "name", "status", "expiration_date", "created_at", "updated_at").
		Values(p.ID, p.PaymentCode, p.Name, p.Status, p.ExpirationDate, p.CreatedAt, p.UpdatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	resp, err := query.RunWith(t.DB).Exec()

	fmt.Print(err)
	fmt.Print(resp)

	return
}

func (t paymentCodeRepository) GetByID(ID string) (res golangtraining.PaymentCode, err error) {
	return
}
