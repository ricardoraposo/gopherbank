package db

import (
	"context"
	"fmt"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
	"github.com/ricardoraposo/gopherbank/models"
)

type DepositRequestDB interface {
	CreateDepositRequest(ctx context.Context, p *models.NewDepositRequestParams) error
	GetRequestsByAccount(ctx context.Context, accountNumber string) ([]*ent.DepositRequest, error)
	GetAllRequests(ctx context.Context) ([]*ent.DepositRequest, error)
	ApproveDepositRequest(ctx context.Context, id int, account string) error
	RejectDepositRequest(ctx context.Context, id int) error
}

type depositRequestDB struct {
    transactionStore TransactionDB
	accountStore AccountDB
	store        *DB
}

func NewDepositRequestDB(store *DB) DepositRequestDB {
	accountDB := NewAccountStore(store)
    transactionDB := NewTransactionDB(store)
	return &depositRequestDB{
		accountStore: accountDB,
		transactionStore: transactionDB,
		store:        store,
	}
}

func (db *depositRequestDB) CreateDepositRequest(ctx context.Context, p *models.NewDepositRequestParams) error {
	acc, err := db.accountStore.GetAccountByNumber(ctx, p.AccountId)
	if err != nil {
		return err
	}

	return db.store.client.DepositRequest.Create().SetAmount(p.Amount).SetAccount(acc).Exec(ctx)
}

func (db *depositRequestDB) GetAllRequests(ctx context.Context) ([]*ent.DepositRequest, error) {
	requests, err := db.store.client.DepositRequest.
		Query().
		WithAccount(func(q *ent.AccountQuery) {
			q.WithUser()
		}).
		Order(ent.Desc(depositrequest.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (db *depositRequestDB) GetRequestsByAccount(ctx context.Context, accountNumber string) ([]*ent.DepositRequest, error) {
	requests, err := db.store.client.DepositRequest.
		Query().
		Where(depositrequest.HasAccountWith(account.ID(accountNumber))).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (db *depositRequestDB) ApproveDepositRequest(ctx context.Context, id int, account string) error {
	deposit, err := db.store.client.DepositRequest.UpdateOneID(id).SetStatus("approved").Save(ctx)
	if err != nil {
		fmt.Println("here")
		return err
	}

    p := &models.DepositParams{
        ToAccountNumber: account,
        Amount: deposit.Amount,
        Type: "deposit",
    }

    if err := db.transactionStore.CreateDepositTransaction(ctx, p); err != nil {
        return err
    }

	return nil
}

func (db *depositRequestDB) RejectDepositRequest(ctx context.Context, id int) error {
	return db.store.client.DepositRequest.UpdateOneID(id).SetStatus("rejected").Exec(ctx)
}
