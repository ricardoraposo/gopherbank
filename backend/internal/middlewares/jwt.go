package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetJWTAccount(c *fiber.Ctx) error {
	claims, ok := c.Context().UserValue("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	return c.JSON(fiber.Map{"number": claims["number"], "admin": claims["admin"]})
}

func JWTAuthentication(c *fiber.Ctx) error {
	authHeader := c.GetReqHeaders()["Authorization"]

	if len(authHeader) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	tokenFields := strings.Fields(authHeader[0])

	if len(tokenFields) != 2 || tokenFields[0] != "Bearer" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	token := tokenFields[1]
	claims, err := getClaimsFromJWT(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	expirationTime := claims["expires"].(float64)
	if int64(expirationTime) < time.Now().Unix() {
		return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
	}

	c.Context().SetUserValue("claims", claims)
	return c.Next()
}

func parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		fmt.Println("invalid signing method", token.Header["alg"])
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret), nil
}

func getClaimsFromJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, parseToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed parsing the token")
	}

	if !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	return claims, nil
}
