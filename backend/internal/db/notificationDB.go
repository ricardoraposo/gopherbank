package db

import (
	"context"

	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/notification"
	"github.com/ricardoraposo/gopherbank/models"
)

type NotificationDB interface {
	CreateNotification(ctx context.Context, p *models.NewNotificationParams) error
    GetNotifications(ctx context.Context, accountNumber string) ([]*ent.Notification, error)
    RemoveNotification(ctx context.Context, id int) error
}

type notificationDB struct {
	accountDB AccountDB
	store     *DB
}

func NewNotificationDB(store *DB) NotificationDB {
	accountDB := NewAccountStore(store)
	return &notificationDB{
		store:     store,
		accountDB: accountDB,
	}
}

func (db *notificationDB) CreateNotification(ctx context.Context, p *models.NewNotificationParams) error {
	acc, err := db.accountDB.GetAccountByNumber(ctx, p.AccountID)
	if err != nil {
		return err
	}

	err = db.store.client.Notification.Create().SetAccount(acc).SetTitle(p.Title).SetContent(p.Content).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *notificationDB) GetNotifications(ctx context.Context, accountNumber string) ([]*ent.Notification, error) {
	acc, err := db.accountDB.GetAccountByNumber(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	return db.store.client.Notification.
        Query().
        Where(notification.HasAccountWith(account.ID(acc.ID))).
        Order(ent.Desc(notification.FieldCreatedAt)).
        All(ctx)
}

func (db *notificationDB) RemoveNotification(ctx context.Context, id int) error {
    return db.store.client.Notification.DeleteOneID(id).Exec(ctx)
}
