package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/ent"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/internal/utils"

	_ "github.com/joho/godotenv/autoload"
)

type AuthParams struct {
	AccountNumber string `json:"number"`
	Password      string `json:"password"`
}

type AuthResponse struct {
	Account *ent.Account `json:"number"`
	Token   string       `json:"token"`
}

type AuthHandler struct {
	store db.AccountDB
}

func NewAuthHandler(client *db.DB) *AuthHandler {
	accountDB := db.NewAccountStore(client)
	return &AuthHandler{store: accountDB}
}

func (h *AuthHandler) Authenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	account, err := h.store.GetAccountByNumber(c.Context(), params.AccountNumber)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Account not found")
	}

	if !utils.ComparePasswords(account.Password, params.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	resp := AuthResponse{
		Account: account,
		Token:   createTokenFromUser(account),
	}

	return c.JSON(resp)
}

func createTokenFromUser(user *ent.Account) string {
	claims := jwt.MapClaims{
		"number":  user.ID,
		"admin":   user.Admin,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Failed to sign token with secret")
	}

	return tokenString
}
