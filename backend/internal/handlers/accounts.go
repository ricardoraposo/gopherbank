package handlers

import (
	"fmt"
	"net/http"

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
	var account models.NewAccountParams
	if err := c.BodyParser(&account); err != nil {
		return err
	}

	num := utils.GenerateAccountNumber()
	encryptedPassword, err := utils.EncryptPassword(account.Password)
	if err != nil {
		return err
	}

	account.Number = num
	account.Password = encryptedPassword

	if err := a.store.CreateAccount(&account); err != nil {
		return err
	}

	return c.JSON(account)
}

func (a *AccountHandler) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return err
	}

	claims := c.Context().Value("claims")
	fmt.Println(claims)

	return c.Status(http.StatusOK).JSON(accounts)
}

func (a *AccountHandler) GetAccountByNumber(c *fiber.Ctx) error {
	number := c.Params("id")
	acc, err := a.store.GetAccountByNumber(number)
	if err != nil {
		return err
	}
	return c.JSON(acc)
}

func (a *AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	number := c.Params("id")
	if err := a.store.DeleteAccount(number); err != nil {
		return err
	}

	return c.JSON(map[string]string{"Message": "Account removed successfully"})
}
