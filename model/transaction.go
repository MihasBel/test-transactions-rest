package model

import (
	"github.com/google/uuid"
	"time"
)

// Transaction DTO model to represent one transaction
type Transaction struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
}
