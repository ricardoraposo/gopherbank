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

func NewAccountHandler(client *db.DB) *AccountHandler {
	accountDB := db.NewAccountStore(client)
	userDB := db.NewUserDB(client)
	return &AccountHandler{
		accountDB: accountDB,
		userDB:    userDB,
	}
}

func (a *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var params models.NewAccountParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	num := utils.GenerateAccountNumber()
	encryptedPassword, err := utils.EncryptPassword(params.Password)
	if err != nil {
		return err
	}

	params.Number = num
	params.Password = encryptedPassword

	account, err := a.accountDB.CreateAccount(c.Context(), &params)
	if err != nil {
		return err
	}

	if err := a.userDB.CreateUser(c.Context(), &params, account); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Account created successfully with number %s", num)})
}

func (a *AccountHandler) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := a.accountDB.GetAllAccounts(c.Context())
	if err != nil {
		return err
	}

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

	account, err := a.accountDB.GetAccountByNumber(c.Context(), number)
	if err != nil {
		return err
	}
	return c.JSON(account)
}

func (a *AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	number := c.Params("id")
	if err := a.userDB.DeleteUser(c.Context(), number); err != nil {
		return err
	}
	if err := a.accountDB.DeleteAccount(c.Context(), number); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"Message": "Account removed successfully"})
}

func (a *AccountHandler) RecoverAccountPassword(c *fiber.Ctx) error {
	params := models.NewPasswordParams{}
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to parse token")
	}

	if claims["number"] != params.Number {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	encryptedPassword, err := utils.EncryptPassword(params.Password)
	if err != nil {
		return err
	}

	if err := a.accountDB.RecoverPassword(c.Context(), encryptedPassword, params.Number); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"Message": "Password updated successfully"})
}
