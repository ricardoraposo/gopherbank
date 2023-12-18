package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	token := c.GetReqHeaders()["Authorization"][0]

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
