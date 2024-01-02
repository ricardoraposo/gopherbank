package db

import (
	"context"
	"fmt"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/user"
	"github.com/ricardoraposo/gopherbank/models"
)

type UserDB interface {
	CreateUser(context.Context, *models.NewAccountParams, *ent.Account) error
	GetUser(context.Context, string) (*ent.User, error)
	EditUser(context.Context, models.EditUserParams, string) error
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
        SetPictureURL(p.PictureURL).
		SetAccount(acc).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	return nil
}

func (u *userDB) GetUser(ctx context.Context, accountNumber string) (*ent.User, error) {
	user, err := u.store.client.User.
		Query().
		Where(user.HasAccountWith(account.ID(accountNumber))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting user: %w", err)
	}

	return user, nil
}

func (u *userDB) EditUser(ctx context.Context, p models.EditUserParams, accountNumber string) error {
	rows, err := u.store.client.User.Update().
		Where(user.HasAccountWith(account.ID(accountNumber))).
		SetFirstName(p.FirstName).
		SetLastName(p.LastName).
		SetEmail(p.Email).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	if rows < 1 {
		return fmt.Errorf("User %s not found", accountNumber)
	}

	return nil
}

func (u *userDB) DeleteUser(ctx context.Context, number string) error {
	rows, err := u.store.client.User.Delete().Where(user.HasAccountWith(account.ID(number))).Exec(ctx)
	if err != nil {
		return err
	}
	if rows < 1 {
		return fmt.Errorf("User %s not found", number)
	}
	return nil
}
