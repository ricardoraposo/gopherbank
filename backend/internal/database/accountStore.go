package database

import "github.com/ricardoraposo/gopherbank/models"

type AccountStore interface {
	CreateAccount(*models.Account) (*models.Account, error)
	GetAllAccounts() ([]*models.Account, error)
}

type accountStore struct {
	store *Store
}

func NewAccountStore(store *Store) AccountStore {
	return &accountStore{store: store}
}

func (a *accountStore) CreateAccount(acc *models.Account) (*models.Account, error) {
	query := "INSERT INTO accounts (first_name, last_name, balance) VALUES (?, ?, ?)"

	_, err := a.store.db.Exec(query, acc.FirstName, acc.LastName, acc.Balance)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (a *accountStore) GetAllAccounts() ([]*models.Account, error) {
	rows, err := a.store.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Balance, &account.CreatedAt, &account.Admin)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}
