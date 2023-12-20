package db

import "github.com/ricardoraposo/gopherbank/models"

type UserDB interface {
	CreateUser(p *models.NewAccountParams) error
}

type userDB struct {
	store *DB
}

func NewUserDB(store *DB) UserDB {
    return &userDB{store: store}
}

func (u *userDB) CreateUser(p *models.NewAccountParams) error {
	query := "INSERT INTO users (first_name, last_name, email, account_number) VALUES (?, ?, ?, ?)"
    _, err := u.store.db.Exec(query, p.FirstName, p.LastName, p.Email, p.Number)
    if err != nil {
        return err
    }

	return nil
}
