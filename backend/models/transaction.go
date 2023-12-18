package models

type Transaction struct {
	ID                int     `json:"id"`
	Amount            float64 `json:"amount"`
	TransferedAt      string  `json:"transferedAt"`
	FromAccountNumber string  `json:"fromAccountNumber"`
	ToAccountNumber   string  `json:"toAccountNumber"`
}

type TransactionParams struct {
	FromAccountNumber string  `json:"fromAccountNumber"`
	ToAccountNumber   string  `json:"toAccountNumber"`
	Amount            float64 `json:"amount"`
}
