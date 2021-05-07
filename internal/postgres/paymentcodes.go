package postgres

import (
	"database/sql"
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

func (t paymentCodeRepository) Create(p *golangtraining.PaymentCode) (res *golangtraining.PaymentCode, err error) {
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
	p.ExpirationDate = now.AddDate(51, 0, 0)

	query := sq.
		Insert("payment_codes").
		Columns("id", "payment_code", "name", "status", "expiration_date", "created_at", "updated_at").
		Values(p.ID, p.PaymentCode, p.Name, p.Status, p.ExpirationDate, p.CreatedAt, p.UpdatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	_, err = query.RunWith(t.DB).Exec()
	if err != nil {
		err = errors.Wrap(err, "error creating data")
	}

	return p, nil
}

func (t paymentCodeRepository) GetByID(ID string) (res golangtraining.PaymentCode, err error) {
	query := sq.
		Select("*").
		Where(sq.Eq{"id": ID}).
		From("payment_codes").
		PlaceholderFormat(sq.Dollar)

	err = query.RunWith(t.DB).QueryRow().Scan(
		&res.ID, &res.PaymentCode, &res.Name, &res.Status, &res.ExpirationDate, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		return res, errors.New("not found")
	}

	return res, nil
}
