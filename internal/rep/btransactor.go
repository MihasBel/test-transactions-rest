package rep

import (
	"context"
	"time"

	"github.com/MihasBel/test-transactions-rest/adapter/broker"
	"github.com/MihasBel/test-transactions-rest/adapter/client/grpc"

	"github.com/go-redis/redis"
	"github.com/ogen-go/ogen/json"

	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// BTransactor transactor with broker
type BTransactor struct {
	b     *broker.Broker
	g     *grpc.Client
	l     zerolog.Logger
	cache *redis.Client
}

// NewBTransactor creates BTransactor
func NewBTransactor(b *broker.Broker, g *grpc.Client, cache *redis.Client, l zerolog.Logger) BTransactor {
	return BTransactor{
		b:     b,
		l:     l,
		g:     g,
		cache: cache,
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

	data, err := json.Marshal(tran)
	if err != nil {
		t.l.Error().Err(err)
	}
	if err = t.cache.Set(tran.ID.String(), string(data), 0).Err(); err != nil {
		t.l.Error().Err(err)
	}
	return tran.ID, nil
}

// GetTransactionByID get transaction info by grpc from service
func (t BTransactor) GetTransactionByID(ctx context.Context, id uuid.UUID) (model.Transaction, error) {
	val, err := t.cache.Get(id.String()).Result()
	if err != nil {
		t.l.Error().Err(err)
	}
	tran := model.Transaction{}
	if err = json.Unmarshal([]byte(val), &tran); err != nil {
		t.l.Error().Err(err)
	}
	if tran.Status != 0 {
		return tran, nil
	}
	rt, err := t.g.GetTransaction(ctx, id)
	if err != nil {
		return model.Transaction{}, err
	}
	data, err := json.Marshal(&rt)
	if err != nil {
		return model.Transaction{}, err
	}
	if err = t.cache.Set(tran.ID.String(), string(data), 0).Err(); err != nil {
		return model.Transaction{}, err
	}
	return *rt, nil

}
