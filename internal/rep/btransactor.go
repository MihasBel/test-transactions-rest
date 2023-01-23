package rep

import (
	"context"
	"fmt"
	"time"

	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/MihasBel/test-transactions-rest/pkg/cache"
	"github.com/go-redis/redis"
	"github.com/ogen-go/ogen/json"

	"github.com/MihasBel/test-transactions-rest/internal/broker"
	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// BTransactor transactor with broker
type BTransactor struct {
	b     *broker.Broker
	cfg   app.Configuration
	l     zerolog.Logger
	cache *redis.Client
}

// NewBTransactor creates BTransactor
func NewBTransactor(b *broker.Broker, l zerolog.Logger, cfg app.Configuration) BTransactor {
	return BTransactor{
		b:     b,
		l:     l,
		cfg:   cfg,
		cache: cache.New(cfg),
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
	pong, err := t.cache.Ping().Result()
	fmt.Println(pong, err)
	data, err := json.Marshal(tran)
	if err != nil {
		t.l.Error().Err(err)
	}
	if err = t.cache.Set(tran.ID.String(), string(data), 0).Err(); err != nil {
		t.l.Error().Err(err)
	}
	return tran.ID, nil
}

// GetTransactionByID TODO connect to service by grpc to get info
func (t BTransactor) GetTransactionByID(_ context.Context, id uuid.UUID) (model.Transaction, error) {
	val, err := t.cache.Get(id.String()).Result()
	if err != nil {
		t.l.Error().Err(err)
	}
	tran := model.Transaction{}
	if err = json.Unmarshal([]byte(val), &tran); err != nil {
		t.l.Error().Err(err)
	}
	return tran, nil
}
