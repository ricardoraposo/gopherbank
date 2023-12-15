package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/database"
	"github.com/ricardoraposo/gopherbank/internal/utils"
	"github.com/ricardoraposo/gopherbank/models"
)

type AccountHandler struct {
	store database.AccountStore
}

func NewAccountHandler(store database.AccountStore) *AccountHandler {
	return &AccountHandler{store: store}
}

func (a *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var account models.Account
	if err := c.BodyParser(&account); err != nil {
		return err
	}

	num := utils.GenerateAccountNumber()
	fmt.Println(num)
	account.Number = num
	acc, err := a.store.CreateAccount(&account)
	if err != nil {
		return err
	}

	account.CreatedAt = time.Now()
	return c.JSON(acc)
}

func (a *AccountHandler) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(accounts)
}
