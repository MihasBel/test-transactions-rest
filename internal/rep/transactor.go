package rep

import (
	"context"
	"time"

	"github.com/MihasBel/test-transactions/model"
)

// Transactor interface to work with transactions repository
type Transactor interface {
	Create(ctx context.Context, uID string, delta int) (model.Transaction, error)
	Balance(ctx context.Context, uID string) (model.Balance, error)
	History(ctx context.Context, uID string) (model.Balance, error)
	ByID(ctx context.Context, tID int) (model.Transaction, error)
	ByTime(ctx context.Context, uID string, tTime time.Time) (model.Transaction, error)
}
