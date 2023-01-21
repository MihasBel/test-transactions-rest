package model

import "time"

// Transaction DTO model to represent one transaction
type Transaction struct {
	ID        int       `json:"id"`
	UserID    string    `json:"userID"`
	BalanceID int       `json:"balanceID"`
	Delta     int       `json:"delta"`
	Time      time.Time `json:"time"`
}

// PostParam DTO model to post one transaction
type PostParam struct {
	UserID string `json:"userID"`
	Delta  int    `json:"delta"`
}
