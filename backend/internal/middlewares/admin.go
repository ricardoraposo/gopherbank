package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAdmin(c *fiber.Ctx) error {
	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
    if !ok {
        return fiber.NewError(fiber.StatusUnauthorized, "Failed to parse token")
    }

    admin := claims["admin"].(bool)

    if !admin {
        return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
    }

	return c.Next()
}
