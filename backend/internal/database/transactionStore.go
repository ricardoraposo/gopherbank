package database

import "github.com/ricardoraposo/gopherbank/models"

type TransactionStore interface{
    CreateTransaction(params *models.TransactionParams) error
}

type transactionStore struct {
	store        *Store
	accountStore AccountStore
}

func NewTransactionStore(store *Store, accountStore AccountStore) TransactionStore {
	return &transactionStore{store, accountStore}
}

func (t *transactionStore) CreateTransaction(params *models.TransactionParams) error {
	query := "INSERT INTO transactions (amount, from_account_number, to_account_number) VALUES (?, ?, ?)"
	if err := t.accountStore.RemoveFromAccount(params.FromAccountNumber, params.Amount); err != nil {
		return err
	}

	if err := t.accountStore.AddToAccount(params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

    _, err := t.store.db.Exec(query, params.Amount, params.FromAccountNumber, params.ToAccountNumber)
	if err != nil {
		return err
	}

	return nil
}
