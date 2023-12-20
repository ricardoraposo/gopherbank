package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/internal/utils"
	"github.com/ricardoraposo/gopherbank/models"
)

type AccountHandler struct {
	accountDB db.AccountDB
	userDB    db.UserDB
}

func NewAccountHandler(accountDB db.AccountDB, userDB db.UserDB) *AccountHandler {
	return &AccountHandler{
		accountDB: accountDB,
		userDB:    userDB,
	}
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

	if err := a.accountDB.CreateAccount(&account); err != nil {
		return err
	}

	if err := a.userDB.CreateUser(&account); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Account created successfully",
	})
}

func (a *AccountHandler) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := a.accountDB.GetAllAccounts()
	if err != nil {
		return err
	}

	claims := c.Context().Value("claims")
	fmt.Println(claims)

	return c.Status(http.StatusOK).JSON(accounts)
}

func (a *AccountHandler) GetAccountByNumber(c *fiber.Ctx) error {
	number := c.Params("id")

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to parse token")
	}

	if claims["number"] != number {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	acc, err := a.accountDB.GetAccountByNumber(number)
	if err != nil {
		return err
	}
	return c.JSON(acc)
}

func (a *AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	number := c.Params("id")
	if err := a.accountDB.DeleteAccount(number); err != nil {
		return err
	}

	return c.JSON(map[string]string{"Message": "Account removed successfully"})
}
