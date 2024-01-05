package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/models"
)

const (
	minNameLength     = 2
	minPasswordLength = 8
)

func ValidateNewAccountParams(c *fiber.Ctx) error {
	var account models.NewAccountParams
	if err := c.BodyParser(&account); err != nil {
		return err
	}
	if len(account.FirstName) < minNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "First name must be at least 2 characters long")
	}
	if len(account.LastName) < minNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "Last name must be at least 2 characters long")
	}
	if len(account.Password) < minPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "Password must have at least 8 characters")
	}

	return c.Next()
}

func ValidateNewPasswordParams(c *fiber.Ctx) error {
	var np models.NewPasswordParams
	if err := c.BodyParser(&np); err != nil {
		return err
	}
	if len(np.Password) < minPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "Password must have at least 8 characters")
	}

	return c.Next()
}

func ValidateLoginParams(c *fiber.Ctx) error {
	var login models.LoginParams
	if err := c.BodyParser(&login); err != nil {
		return err
	}

	if len(login.Number) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Account number must be provided")
	}

	if len(login.Password) < minPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "Password must have at least 8 characters")
	}

	return c.Next()
}

func ValidateTransferParams(c *fiber.Ctx) error {
	var transfer models.TransferParams
	if err := c.BodyParser(&transfer); err != nil {
		return err
	}

	if len(transfer.ToAccountNumber) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Destiny account number must be provided")
	}

	if len(transfer.FromAccountNumber) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Origin account number must be provided")
	}

	if transfer.Amount <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid amount")
	}

	return c.Next()
}

func ValidateWithdrawParams(c *fiber.Ctx) error {
	var withdraw models.WithdrawParams
	if err := c.BodyParser(&withdraw); err != nil {
		return err
	}

	if len(withdraw.FromAccountNumber) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Origin account number must be provided")
	}

	if withdraw.Amount <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid amount")
	}

	return c.Next()
}
