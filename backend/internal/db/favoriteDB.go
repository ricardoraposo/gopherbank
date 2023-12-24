package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/models"
)

type FavoriteDB interface {
	CreateFavorite(context.Context, models.NewFavoriteParams) error
}

type favoriteDB struct {
	db        *DB
	accountDB AccountDB
}

func NewFavoriteDB(client *DB) FavoriteDB {
	accountDB := NewAccountStore(client)
	return FavoriteDB(&favoriteDB{client, accountDB})
}

func (f *favoriteDB) CreateFavorite(ctx context.Context, p models.NewFavoriteParams) error {

	account, err := f.accountDB.GetAccountByNumber(ctx, p.AccountID)
	if err != nil {
		return err
	}

	favorited, err := f.accountDB.GetAccountByNumber(ctx, p.FavoritedID)
	if err != nil {
		return err
	}

    return account.Update().AddFavorites(favorited).Exec(ctx)
}
