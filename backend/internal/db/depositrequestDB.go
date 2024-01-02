package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
	"github.com/ricardoraposo/gopherbank/models"
)

type DepositRequestDB interface {
	CreateDepositRequest(ctx context.Context, p *models.NewDepositRequestParams) error
	GetRequestsByAccount(ctx context.Context, accountNumber string) ([]*ent.DepositRequest, error)
    GetAllRequests(ctx context.Context) ([]*ent.DepositRequest, error)
}

type depositRequestDB struct {
	accountStore AccountDB
	store        *DB
}

func NewDepositRequestDB(store *DB) DepositRequestDB {
	accountDB := NewAccountStore(store)
	return &depositRequestDB{
		accountStore: accountDB,
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
		WithAccount().
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
