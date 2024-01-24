package db

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	c "github.com/ricardoraposo/gopherbank/config"
	"github.com/ricardoraposo/gopherbank/ent"
)

type DB struct {
	client *ent.Client
}

func New() *DB {
	client, err := ent.Open("mysql", c.ConnString())
	if err != nil {
		log.Fatal(err)
	}

    ctx := context.Background()

    client.Schema.Create(ctx)

	return &DB{client}
}
