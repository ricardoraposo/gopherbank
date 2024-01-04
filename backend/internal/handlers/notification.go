package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/db"
)

type NotificationHandler struct {
	store db.NotificationDB
}

func NewNotificationHandler(client *db.DB) *NotificationHandler {
	notificationDB := db.NewNotificationDB(client)
	return &NotificationHandler{
		store: notificationDB,
	}
}

func (h *NotificationHandler) GetAccountNotifications(c *fiber.Ctx) error {
	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		fmt.Println("could not get the claims")
		return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
	}

	accountNumber, ok := claims["number"].(string)

	notifications, err := h.store.GetNotifications(c.Context(), accountNumber)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"notifications": notifications})
}

func (h *NotificationHandler) RemoveNotification(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid id")
    }

    if err := h.store.RemoveNotification(c.Context(), id); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    return c.JSON(fiber.Map{"message": "notification removed"})
}
