package database

import "github.com/ricardoraposo/gopherbank/models"

type TransactionStore interface {
	CreateTransferTransaction(params *models.TransferParams) error
	CreateDepositTransaction(params *models.DepositParams) error
    CreateWithdrawTransaction(params *models.WithdrawParams) error
}

type transactionStore struct {
	store        *Store
	accountStore AccountStore
}

func NewTransactionStore(store *Store, accountStore AccountStore) TransactionStore {
	return &transactionStore{store, accountStore}
}

func (t *transactionStore) CreateTransferTransaction(params *models.TransferParams) error {
	if err := t.accountStore.Transfer(params.FromAccountNumber, params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (amount, from_account_number, to_account_number) VALUES (?, ?, ?)"
	_, err := t.store.db.Exec(query, params.Amount, params.FromAccountNumber, params.ToAccountNumber)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionStore) CreateDepositTransaction(params *models.DepositParams) error {
	if err := t.accountStore.AddToAccount(params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (amount, to_account_number) VALUES (?, ?)"
	_, err := t.store.db.Exec(query, params.Amount, params.ToAccountNumber)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionStore) CreateWithdrawTransaction(params *models.WithdrawParams) error {
	if err := t.accountStore.RemoveFromAccount(params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (amount, to_account_number) VALUES (?, ?)"
	_, err := t.store.db.Exec(query, params.Amount, params.ToAccountNumber)
	if err != nil {
		return err
	}

	return nil
}
