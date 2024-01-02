package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/models"
)

type DepositRequestHandler struct {
	db db.DepositRequestDB
}

func NewDepositRequestHandler(client *db.DB) *DepositRequestHandler {
	depositRequestDB := db.NewDepositRequestDB(client)
	return &DepositRequestHandler{
		db: depositRequestDB,
	}
}

func (h *DepositRequestHandler) CreateDepositRequest(c *fiber.Ctx) error {
	params := &models.NewDepositRequestParams{}
	if err := c.BodyParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "error parsing token")
	}

	number, ok := claims["number"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "error getting account number")
	}

	params.AccountId = number

	if err := h.db.CreateDepositRequest(c.Context(), params); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintln("error creating deposit request:", err))
	}

	return c.JSON(fiber.Map{"message": "deposit request created"})
}

func (h *DepositRequestHandler) GetAllDepositRequests(c *fiber.Ctx) error {
	requests, err := h.db.GetAllRequests(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error getting deposit requests")
	}

	return c.JSON(fiber.Map{"requests": requests})
}

func (h *DepositRequestHandler) GetDepositRequestsByAccount(c *fiber.Ctx) error {
	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "error parsing token")
	}

	accountNumber, ok := claims["number"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "error getting account number")
	}

	requests, err := h.db.GetRequestsByAccount(c.Context(), accountNumber)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error getting deposit requests")
	}

	return c.JSON(fiber.Map{"requests": requests})
}

func (h *DepositRequestHandler) ApproveRequest(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	p := &models.DepositParam{}
	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
    fmt.Println(p)

	if err := h.db.ApproveDepositRequest(c.Context(), id, p.Account); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintln("error approving deposit request:", err))
	}

	return c.JSON(fiber.Map{"message": "deposit request approved"})
}

func (h *DepositRequestHandler) RejectRequest(c *fiber.Ctx) error {
	idString := c.Params("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.db.RejectDepositRequest(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error rejecting deposit request")
	}

	return c.JSON(fiber.Map{"message": "deposit request rejected"})
}
