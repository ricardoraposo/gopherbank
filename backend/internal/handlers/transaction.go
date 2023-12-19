package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/database"
	"github.com/ricardoraposo/gopherbank/models"
)

type TransactionHandler struct {
	store database.TransactionStore
}

func NewTransactionHandler(store database.TransactionStore) *TransactionHandler {
	return &TransactionHandler{store: store}
}

func (h *TransactionHandler) Transfer(c *fiber.Ctx) error {
	params := &models.TransferParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to parse token")
	}

    fmt.Println("claims", claims)
	if claims["number"] != params.FromAccountNumber {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	if err := h.store.CreateTransferTransaction(params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Transfer ocurred successfully"})
}

func (h *TransactionHandler) Deposit(c *fiber.Ctx) error {
	params := &models.DepositParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := h.store.CreateDepositTransaction(params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Deposit ocurred successfully"})
}

func (h *TransactionHandler) Withdraw(c *fiber.Ctx) error {
	params := &models.WithdrawParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to parse token")
	}

	if claims["number"] != params.FromAccountNumber {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	if err := h.store.CreateWithdrawTransaction(params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Withdraw ocurred successfully"})
}
