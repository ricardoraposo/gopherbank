package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/database"
	"github.com/ricardoraposo/gopherbank/models"
)

type TransactionHandler struct {
	store database.TransactionStore
}

func NewTransactionHandler(store database.TransactionStore) *TransactionHandler {
	return &TransactionHandler{store: store}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	params := &models.TransactionParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := h.store.CreateTransaction(params); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Transaction ocurred successfully"})
}
