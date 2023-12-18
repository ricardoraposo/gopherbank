package database

import (
	"database/sql"
	"fmt"

	"github.com/ricardoraposo/gopherbank/models"
)

type AccountStore interface {
	CreateAccount(*models.NewAccountParams) error
	GetAllAccounts() ([]*models.Account, error)
	GetAccountByNumber(number string) (*models.Account, error)
	DeleteAccount(number string) error
}

type accountStore struct {
	store *Store
}

func NewAccountStore(store *Store) AccountStore {
	return &accountStore{store: store}
}

func (a *accountStore) CreateAccount(acc *models.NewAccountParams) error {
	query := `INSERT INTO accounts (first_name, last_name, password, number)
    VALUES (?, ?, ?, ?)`

	_, err := a.store.db.Exec(query, acc.FirstName, acc.LastName, acc.Password, acc.Number)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountStore) GetAccountByNumber(number string) (*models.Account, error) {
	res, err := a.store.db.Query("SELECT * FROM accounts WHERE number = ?", number)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		return scanRow(res)
	}

	return nil, fmt.Errorf(fmt.Sprintf("Account number %s not found", number))
}

func (a *accountStore) GetAllAccounts() ([]*models.Account, error) {
	rows, err := a.store.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}
	for rows.Next() {
		account, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (a *accountStore) DeleteAccount(number string) error {
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

func scanRow(row *sql.Rows) (*models.Account, error) {
	var account models.Account
	err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Password, &account.Balance, &account.CreatedAt, &account.Admin)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
