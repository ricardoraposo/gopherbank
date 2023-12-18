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
