package models

type TransactionDetails struct {
	TransactionID int     `json:"transactionID"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
	TransactedAt  string  `json:"transactedAt"`
}
