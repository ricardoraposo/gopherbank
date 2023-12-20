package db

import (
	"database/sql"
	"fmt"

	"github.com/ricardoraposo/gopherbank/internal/utils"
	"github.com/ricardoraposo/gopherbank/models"
)

type AccountDB interface {
	CreateAccount(*models.NewAccountParams) error
	GetAllAccounts() ([]*models.DisplayAccount, error)
	GetAccountByNumber(number string) (*models.DisplayAccount, error)
	DeleteAccount(number string) error
	AddToAccount(number string, amount float64) error
	RemoveFromAccount(number string, amount float64) error
	Transfer(from, to string, amount float64) error
}

type accountDB struct {
	store *DB
}

func NewAccountStore(store *DB) AccountDB {
	return &accountDB{store: store}
}

func (a *accountDB) Transfer(from, to string, amount float64) error {
	tx, err := a.store.db.Begin()
	if err != nil {
		return err
	}

	err = a.RemoveFromAccount(from, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = a.AddToAccount(to, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (a *accountDB) AddToAccount(number string, amount float64) error {
	query := `UPDATE accounts SET balance = balance + ? WHERE number = ?`
	res, err := a.store.db.Exec(query, amount, number)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

func (a *accountDB) RemoveFromAccount(number string, amount float64) error {
	query := `UPDATE accounts SET balance = balance - ? WHERE number = ?`
	res, err := a.store.db.Exec(query, amount, number)
	if err != nil {
		return fmt.Errorf("Not enough funds to conclude transaction")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

func (a *accountDB) CreateAccount(acc *models.NewAccountParams) error {
	query := `INSERT INTO accounts (password, number)
    VALUES (?, ?)`

	_, err := a.store.db.Exec(query, acc.Password, acc.Number)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountDB) GetAccountByNumber(number string) (*models.DisplayAccount, error) {
	res, err := a.store.db.Query(utils.GetAccountByNumberQuery, number)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		return scanRow(res)
	}

	return nil, fmt.Errorf(fmt.Sprintf("Account number %s not found", number))
}

func (a *accountDB) GetAllAccounts() ([]*models.DisplayAccount, error) {
	rows, err := a.store.db.Query(utils.GetAllAccountsQuery)
	if err != nil {
		return nil, err
	}

	accounts := []*models.DisplayAccount{}
	for rows.Next() {
		account, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (a *accountDB) DeleteAccount(number string) error {
	res, err := a.store.db.Exec("DELETE FROM accounts WHERE number = ?", number)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows < 1 {
		return fmt.Errorf("Account %s not found", number)
	}

	return nil
}

func scanRow(row *sql.Rows) (*models.DisplayAccount, error) {
	var account models.DisplayAccount
	err := row.Scan(&account.FirstName, &account.LastName, &account.Email, &account.Number, &account.Balance, &account.Password, &account.Admin)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
