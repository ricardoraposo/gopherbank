package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/models"
)

type FavoriteDB interface {
	ToggleFavorite(context.Context, models.NewFavoriteParams) error
	GetFavoritedsByAccount(context.Context, string) ([]*ent.Account, error)
}

type favoriteDB struct {
	db        *DB
	accountDB AccountDB
}

func NewFavoriteDB(client *DB) FavoriteDB {
	accountDB := NewAccountStore(client)
	return FavoriteDB(&favoriteDB{client, accountDB})
}

func (f *favoriteDB) ToggleFavorite(ctx context.Context, p models.NewFavoriteParams) error {
	acc, err := f.accountDB.GetAccountByNumber(ctx, p.AccountID)
	if err != nil {
		return err
	}

	hasFavorite, err := acc.QueryFavorites().Where(account.ID(p.FavoritedID)).Exist(ctx)
	if err != nil {
		return err
	}

	favorited, err := f.accountDB.GetAccountByNumber(ctx, p.FavoritedID)
	if err != nil {
		return err
	}

	if hasFavorite {
		return acc.Update().RemoveFavorites(favorited).Exec(ctx)
	} else {
		return acc.Update().AddFavorites(favorited).Exec(ctx)
	}
}

func (f *favoriteDB) GetFavoritedsByAccount(ctx context.Context, accountNumber string) ([]*ent.Account, error) {
	account, err := f.db.client.Account.
		Query().
		Where(account.ID(accountNumber)).
		WithFavorites(func (q *ent.AccountQuery) {
            q.WithUser()
        }).
		Only(ctx)

	favorites := account.Edges.Favorites

	if err != nil {
		return nil, err
	}

	return favorites, nil
}
