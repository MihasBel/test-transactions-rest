package mocks

import (
	"context"
	"time"

	"github.com/MihasBel/test-transactions/model"
	"github.com/rs/zerolog/log"
)

// TestTransactor mock realization of transactor interface
type TestTransactor struct {
}

// Create post a transaction
func (t TestTransactor) Create(_ context.Context, uID string, delta int) (model.Transaction, error) {
	log.Info().Msgf("Create was called for userID: %v and delta: %v", uID, delta)
	return model.Transaction{
		ID:        1,
		UserID:    "uID",
		BalanceID: 1,
		Delta:     100,
		Time:      time.Now(),
	}, nil
}

// Balance return a current balance of user
func (t TestTransactor) Balance(_ context.Context, uID string) (model.Balance, error) {
	log.Info().Msgf("Balance was called for userID: %v", uID)
	return model.Balance{}, nil
}

// History return a current balance with all related transactions of user
func (t TestTransactor) History(_ context.Context, uID string) (model.Balance, error) {
	log.Info().Msgf("History was called for userID: %v", uID)
	val := []model.Transaction{{
		ID:        1,
		UserID:    uID,
		BalanceID: 1,
		Delta:     100,
		Time:      time.Now(),
	}, {
		ID:        2,
		UserID:    uID,
		BalanceID: 1,
		Delta:     -20,
		Time:      time.Now(),
	},
	}
	bal := model.Balance{
		ID:          1,
		UserID:      uID,
		Amount:      80,
		TranHistory: val,
	}
	return bal, nil
}

// ByID return specific transaction by transaction ID
func (t TestTransactor) ByID(_ context.Context, tID int) (model.Transaction, error) {
	log.Info().Msgf("ByID was called for transactionID: %v", tID)
	return model.Transaction{
		ID:        1,
		UserID:    "uId",
		BalanceID: 1,
		Delta:     100,
		Time:      time.Now(),
	}, nil
}

// ByTime return specific transaction related to user by time
func (t TestTransactor) ByTime(_ context.Context, uID string, tTime time.Time) (model.Transaction, error) {
	log.Info().Msgf("ByTime was called for userID: %v and time: $v", uID, tTime)
	return model.Transaction{
		ID:        1,
		UserID:    uID,
		BalanceID: 1,
		Delta:     100,
		Time:      tTime,
	}, nil
}
