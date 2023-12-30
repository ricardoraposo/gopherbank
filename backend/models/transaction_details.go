package models

import "github.com/ricardoraposo/gopherbank/ent"

type TransactionDetails struct {
	TransactionID int     `json:"transactionID"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
	TransactedAt  string  `json:"transactedAt"`
}
type ByTransactedAt []*ent.Transaction

func (a ByTransactedAt) Len() int {
	return len(a)
}

func (a ByTransactedAt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTransactedAt) Less(i, j int) bool {
	return a[i].Edges.Detail.TransactedAt.After(a[j].Edges.Detail.TransactedAt)
}
