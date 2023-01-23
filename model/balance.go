package model

import (
	"time"

	"github.com/google/uuid"
)

// Balance represent slice of balance of user.
type Balance struct {
	ID     int       `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Amount int       `json:"amount"`
	Month  time.Time `json:"month"`
}
