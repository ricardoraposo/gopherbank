package main

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	c "github.com/ricardoraposo/gopherbank/config"
	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
)

func main() {
	client, err := ent.Open("mysql", c.ConnString())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	deposits := client.DepositRequest.
		Query().
		Where(depositrequest.HasAccountWith(account.ID("00276710"))).
		AllX(ctx)

	dj, _ := json.Marshal(deposits)

	requests, err := client.DepositRequest.
		Query().
		WithAccount().
		First(ctx)

	rj, _ := json.Marshal(requests)

	fmt.Println(string(dj))
	fmt.Println(string(rj))
}
