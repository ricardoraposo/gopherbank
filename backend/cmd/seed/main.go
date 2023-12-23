package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	c "github.com/ricardoraposo/gopherbank/config"
	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/internal/utils"
)

func main() {
	client, err := ent.Open("mysql", c.ConnString())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	client.User.Delete().ExecX(ctx)
	client.Account.Delete().ExecX(ctx)

	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	for _, p := range Params {
		acc := createAccount(client, ctx, p)
		createUser(client, ctx, p, acc)
	}

	createAdmin(client, ctx)
}

func createAccount(client *ent.Client, ctx context.Context, p NewAccountParams) *ent.Account {
	encrpytedPwd, err := utils.EncryptPassword(p.Password)
	if err != nil {
		panic(err)
	}

	return client.Account.Create().SetID(p.Number).SetPassword(encrpytedPwd).SetBalance(100).SaveX(ctx)
}

func createUser(client *ent.Client, ctx context.Context, p NewAccountParams, acc *ent.Account) {
	client.User.Create().
		SetFirstName(p.FirstName).
		SetLastName(p.LastName).
		SetEmail(p.Email).
		SetAccount(acc).
		ExecX(ctx)
}

func createAdmin(client *ent.Client, ctx context.Context) {
	encrpytedPwd, err := utils.EncryptPassword("eusouoadmin")
	if err != nil {
		panic(err)
	}
	num := utils.GenerateAccountNumber()

	acc := client.Account.Create().SetID(num).SetPassword(encrpytedPwd).SetBalance(100).SetAdmin(true).SaveX(ctx)
	client.User.Create().
		SetFirstName("Rick").
		SetLastName("Raposo").
		SetEmail("admin@gopher.com").
		SetAccount(acc).
		ExecX(ctx)
}

type NewAccountParams struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
	Number    string
}

var Params = []NewAccountParams{
	NewAccountParams{
		FirstName: "Ricardo",
		LastName:  "Gopher",
		Password:  "123456789",
		Email:     "rick@Gopher.com",
		Number:    utils.GenerateAccountNumber(),
	},
	NewAccountParams{
		FirstName: "Xico",
		LastName:  "Gopher",
		Password:  "123456789",
		Email:     "xico@gopher.com",
		Number:    utils.GenerateAccountNumber(),
	},
	NewAccountParams{
		FirstName: "Willy",
		LastName:  "Gopher",
		Password:  "123456789",
		Email:     "willy@gopher.com",
		Number:    utils.GenerateAccountNumber(),
	},
	NewAccountParams{
		FirstName: "Tito",
		LastName:  "Gopher",
		Password:  "123456789",
		Email:     "tito@gopher.com",
		Number:    utils.GenerateAccountNumber(),
	},
}
