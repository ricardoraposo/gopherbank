package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/models"
)

type TransactionDB interface {
	CreateTransferTransaction(context.Context, *models.TransferParams) error
	CreateDepositTransaction(context.Context, *models.DepositParams) error
	CreateWithdrawTransaction(context.Context, *models.WithdrawParams) error
}

type transactionDB struct {
	store        *DB
	accountStore AccountDB
}

func NewTransactionDB(client *DB) TransactionDB {
	accountDB := NewAccountStore(client)
	return &transactionDB{client, accountDB}
}

func (t *transactionDB) CreateTransferTransaction(ctx context.Context, params *models.TransferParams) error {
	if err := t.accountStore.Transfer(ctx, params.FromAccountNumber, params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	fromAccount, err := t.store.client.Account.Get(ctx, params.FromAccountNumber)
	if err != nil {
		return err
	}
	toAccount, err := t.store.client.Account.Get(ctx, params.ToAccountNumber)
	if err != nil {
		return err
	}
	transaction, err := t.store.client.Transaction.Create().SetFromAccount(fromAccount).SetToAccount(toAccount).Save(ctx)
	if err != nil {
		return err
	}

	err = t.store.client.TransactionDetail.Create().SetTransactionID(transaction.ID).SetAmount(params.Amount).SetType(params.Type).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionDB) CreateDepositTransaction(ctx context.Context, params *models.DepositParams) error {
	if err := t.accountStore.AddToAccount(ctx, params.ToAccountNumber, params.Amount); err != nil {
		return err
	}

	toAccount, err := t.store.client.Account.Get(ctx, params.ToAccountNumber)
	if err != nil {
		return err
	}
	transaction, err := t.store.client.Transaction.Create().SetToAccount(toAccount).Save(ctx)
	if err != nil {
		return err
	}

	err = t.store.client.TransactionDetail.Create().SetTransactionID(transaction.ID).SetAmount(params.Amount).SetType(params.Type).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionDB) CreateWithdrawTransaction(ctx context.Context, params *models.WithdrawParams) error {
	if err := t.accountStore.RemoveFromAccount(ctx, params.FromAccountNumber, params.Amount); err != nil {
		return err
	}

	fromAccount, err := t.store.client.Account.Get(ctx, params.FromAccountNumber)
	if err != nil {
		return err
	}

	transaction, err := t.store.client.Transaction.Create().SetFromAccount(fromAccount).Save(ctx)
	if err != nil {
		return err
	}

	err = t.store.client.TransactionDetail.Create().SetTransactionID(transaction.ID).SetAmount(params.Amount).SetType(params.Type).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
