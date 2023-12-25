package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/transaction"
	"github.com/ricardoraposo/gopherbank/models"
)

type TransactionDB interface {
	GetAllTransactions(context.Context) ([]*ent.TransactionDetails, error)
	GetAccountTransactions(context.Context, string) ([]*ent.TransactionDetails, error)
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

	transaction, err := t.store.client.Transaction.
		Create().
		SetToAccount(toAccount).
		SetFromAccount(fromAccount).Save(ctx)

	return t.store.client.TransactionDetails.Create().
		SetType(params.Type).
		SetAmount(params.Amount).
		SetTransaction(transaction).
		Exec(ctx)
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

	return t.store.client.TransactionDetails.Create().
		SetType(params.Type).
		SetAmount(params.Amount).
		SetTransaction(transaction).
		Exec(ctx)
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

	return t.store.client.TransactionDetails.Create().
		SetType(params.Type).
		SetAmount(params.Amount).
		SetTransaction(transaction).
		Exec(ctx)
}

func (t *transactionDB) GetAllTransactions(ctx context.Context) ([]*ent.TransactionDetails, error) {
	transactions, err := t.store.client.TransactionDetails.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionDB) GetAccountTransactions(ctx context.Context, accountNumber string) ([]*ent.TransactionDetails, error) {
	toTransactions, err := t.store.client.Transaction.
        Query().
        Where(transaction.HasToAccountWith(account.ID(accountNumber))).
        QueryDetail().
        All(ctx)

	fromTransactions, err := t.store.client.Transaction.
        Query().
        Where(transaction.HasFromAccountWith(account.ID(accountNumber))).
        QueryDetail().
        All(ctx)

    for _, transaction := range fromTransactions {
        transaction.Amount = -transaction.Amount
    }

    toTransactions = append(toTransactions, fromTransactions...)

    if err != nil {
        return nil, err
    }

	return toTransactions, nil
}
