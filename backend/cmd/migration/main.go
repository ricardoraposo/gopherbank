package main

import (
	"context"

	"github.com/ricardoraposo/gopherbank/ent"
	c "github.com/ricardoraposo/gopherbank/config"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", c.ConnString())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
}
