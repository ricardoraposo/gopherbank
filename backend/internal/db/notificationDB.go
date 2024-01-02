package db

type NotificationDB interface {}

type notificationDB struct {
    store *DB
}

func NewNotificationDB(store *DB) NotificationDB {
    return &notificationDB{
        store: store,
    }
}

func (db *notificationDB) CreateNotification() {
    db.store.client.Notification.Create()
}
