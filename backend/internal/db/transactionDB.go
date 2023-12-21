package db

import "github.com/ricardoraposo/gopherbank/models"

type TransactionDB interface {
	CreateTransferTransaction(params *models.TransferParams) error
	CreateDepositTransaction(params *models.DepositParams) error
	CreateWithdrawTransaction(params *models.WithdrawParams) error
}

type transactionDB struct {
	store        *DB
	accountStore AccountDB
}

func NewTransactionDB(client *DB) TransactionDB {
    accountDB := NewAccountStore(client)
	return &transactionDB{client, accountDB}
}

func (t *transactionDB) CreateTransferTransaction(params *models.TransferParams) error {
	if err := t.accountStore.Transfer(params.FromAccountNumber, params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (from_account_number, to_account_number) VALUES (?, ?)"
	res, err := t.store.db.Exec(query, params.FromAccountNumber, params.ToAccountNumber)
	if err != nil {
		return err
	}

	query = "INSERT INTO transaction_details (transaction_id, amount, type) VALUES (?, ?, ?)"
	transactionID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = t.store.db.Exec(query, transactionID, params.Amount, params.Type)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionDB) CreateDepositTransaction(params *models.DepositParams) error {
	if err := t.accountStore.AddToAccount(params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (to_account_number) VALUES (?)"
	res, err := t.store.db.Exec(query, params.ToAccountNumber)
	if err != nil {
		return err
	}

	transactionID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	query = "INSERT INTO transaction_details (transaction_id, amount, type) VALUES (?, ?, ?)"
	_, err = t.store.db.Exec(query, transactionID, params.Amount, params.Type)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionDB) CreateWithdrawTransaction(params *models.WithdrawParams) error {
	if err := t.accountStore.RemoveFromAccount(params.FromAccountNumber, params.Amount); err != nil {
		return err
	}

	query := "INSERT INTO transactions (from_account_number) VALUES (?)"
	res, err := t.store.db.Exec(query, params.FromAccountNumber)
	if err != nil {
		return err
	}

	transactionID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	query = "INSERT INTO transaction_details (transaction_id, amount, type) VALUES (?, ?, ?)"
	_, err = t.store.db.Exec(query, transactionID, params.Amount, params.Type)
	if err != nil {
		return err
	}

	return nil
}
