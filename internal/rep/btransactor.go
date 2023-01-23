package rep

import (
	"context"
	"time"

	"github.com/MihasBel/test-transactions-rest/internal/broker"
	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// BTransactor transactor with broker
type BTransactor struct {
	b *broker.Broker
	l zerolog.Logger
}

// NewBTransactor creates BTransactor
func NewBTransactor(b *broker.Broker, l zerolog.Logger) BTransactor {
	return BTransactor{
		b: b,
		l: l,
	}
}

// PlaceTransaction post transaction to broker
func (t BTransactor) PlaceTransaction(ctx context.Context, amount int, userID uuid.UUID) (uuid.UUID, error) {
	tran := model.Transaction{
		ID:          uuid.New(),
		UserID:      userID,
		Amount:      amount,
		CreatedAt:   time.Now(),
		Status:      0,
		Description: "in processing",
	}
	if err := t.b.PlaceTransaction(ctx, tran); err != nil {
		return [16]byte{}, err
	}
	return tran.ID, nil
}

// GetTransactionByID TODO connect to service by grpc to get info
func (t BTransactor) GetTransactionByID(_ context.Context, _ uuid.UUID) (model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
