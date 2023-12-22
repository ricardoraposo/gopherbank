package db

import (
	"context"
	"fmt"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/user"
	"github.com/ricardoraposo/gopherbank/models"
)

type UserDB interface {
	CreateUser(context.Context, *models.NewAccountParams, *ent.Account) error
	DeleteUser(context.Context, string) error
}

type userDB struct {
	store *DB
}

func NewUserDB(store *DB) UserDB {
	return &userDB{store: store}
}

func (u *userDB) CreateUser(ctx context.Context, p *models.NewAccountParams, acc *ent.Account) error {
	err := u.store.client.User.
		Create().
		SetEmail(p.Email).
		SetFirstName(p.FirstName).
		SetLastName(p.LastName).
		SetAccount(acc).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	return nil
}

func (u *userDB) DeleteUser(ctx context.Context, number string) error {
	rows, err := u.store.client.User.Delete().Where(user.ID(number)).Exec(ctx)
	if err != nil {
		return err
	}
	if rows < 1 {
		return fmt.Errorf("User %s not found", number)
	}
	return nil
}
