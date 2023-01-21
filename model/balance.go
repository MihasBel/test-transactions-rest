package model

// Balance represent balance of user. can contain related transactions.
type Balance struct {
	ID          int           `json:"id"`
	UserID      string        `json:"userID"`
	Amount      int           `json:"amount"`
	TranHistory []Transaction `json:"tranHistory"`
}
