package rep

import (
	"context"

	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/google/uuid"
)

//go:generate mockgen -source=transactor.go -destination=../../mocks/test-transactor.go -package=mocks

// Transactor interface to work with transactions
type Transactor interface {
	PlaceTransaction(ctx context.Context, amount int, userID uuid.UUID) (uuid.UUID, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (model.Transaction, error)
}
