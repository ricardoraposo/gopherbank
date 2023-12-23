package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/models"
)

type TransactionHandler struct {
	store db.TransactionDB
}

func NewTransactionHandler(client *db.DB) *TransactionHandler {
	transactionDB := db.NewTransactionDB(client)
	return &TransactionHandler{store: transactionDB}
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

	if claims["number"] != params.FromAccountNumber {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	params.Type = "transfer"
	if err := h.store.CreateTransferTransaction(c.Context(), params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Transfer ocurred successfully"})
}

func (h *TransactionHandler) Deposit(c *fiber.Ctx) error {
	params := &models.DepositParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	params.Type = "deposit"
	if err := h.store.CreateDepositTransaction(c.Context(), params); err != nil {
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

	params.Type = "withdraw"
	if err := h.store.CreateWithdrawTransaction(c.Context(), params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Withdraw ocurred successfully"})
}
