package db

import (
	"context"
	"fmt"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/models"
)

type AccountDB interface {
	CreateAccount(context.Context, *models.NewAccountParams) (*ent.Account, error)
	GetAllAccounts(context.Context) ([]*ent.Account, error)
	GetAccountByNumber(context.Context, string) (*ent.Account, error)
	DeleteAccount(context.Context, string) error
    AddToAccount(ctx context.Context, to string, amount float64) error
	RemoveFromAccount(ctx context.Context, from string, amount float64) error
	Transfer(ctx context.Context, from string, to string, amount float64) error
}

type accountDB struct {
	store *DB
}

func NewAccountStore(store *DB) AccountDB {
	return &accountDB{store: store}
}

func (a *accountDB) Transfer(ctx context.Context, from, to string, amount float64) error {
	tx, err := a.store.client.Tx(ctx)
	if err != nil {
		return err
	}

	err = a.RemoveFromAccount(ctx, from, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = a.AddToAccount(ctx, to, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (a *accountDB) AddToAccount(ctx context.Context, number string, amount float64) error {
	rows, err := a.store.client.Account.Update().Where(account.ID(number)).AddBalance(amount).Save(ctx)
    fmt.Println(rows)
	if err != nil {
		return err
	}

	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

func (a *accountDB) RemoveFromAccount(ctx context.Context, number string, amount float64) error {
	rows, err := a.store.client.Account.Update().Where(account.ID(number)).AddBalance(-amount).Save(ctx)
	if err != nil {
		return fmt.Errorf("Not enough funds to conclude transaction")
	}

	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

func (a *accountDB) CreateAccount(ctx context.Context, acc *models.NewAccountParams) (*ent.Account, error) {
	account, err := a.store.client.Account.Create().SetID(acc.Number).SetPassword(acc.Password).Save(ctx)
	if err != nil {
		return &ent.Account{}, err
	}

	return account, nil
}

func (a *accountDB) GetAccountByNumber(ctx context.Context, number string) (*ent.Account, error) {
	account, err := a.store.client.Account.Query().Where(account.ID(number)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *accountDB) GetAllAccounts(ctx context.Context) ([]*ent.Account, error) {
	accounts, err := a.store.client.Account.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *accountDB) DeleteAccount(ctx context.Context, number string) error {
	rows, err := a.store.client.Account.Delete().Where(account.ID(number)).Exec(ctx)
	if err != nil {
		return err
	}

	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

// func scanRow(row *sql.Rows) (*models.DisplayAccount, error) {
// 	var account models.DisplayAccount
// 	err := row.Scan(&account.FirstName, &account.LastName, &account.Email, &account.Number, &account.Balance, &account.Password, &account.Admin)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &account, nil
// }