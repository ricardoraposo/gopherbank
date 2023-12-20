package models

type Transaction struct {
	ID                int    `json:"id"`
	FromAccountNumber string `json:"fromAccountNumber"`
	ToAccountNumber   string `json:"toAccountNumber"`
}

type TransferParams struct {
	FromAccountNumber string  `json:"fromAccountNumber"`
	ToAccountNumber   string  `json:"toAccountNumber"`
	Type              string  `json:"type"`
	Amount            float64 `json:"amount"`
}

type DepositParams struct {
	ToAccountNumber string  `json:"toAccountNumber"`
	Amount          float64 `json:"amount"`
	Type            string  `json:"type"`
}

type WithdrawParams struct {
	FromAccountNumber string  `json:"fromAccountNumber"`
	Amount            float64 `json:"amount"`
	Type              string  `json:"type"`
}
