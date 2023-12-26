package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/models"
)

type UserHandler struct {
	userDB db.UserDB
}

func NewUserHandler(client *db.DB) *UserHandler {
	userDB := db.NewUserDB(client)
	return &UserHandler{userDB}
}

func (u *UserHandler) EditUser(c *fiber.Ctx) error {
    ctx := c.Context()
	accountId := c.Params("id")
	params := models.EditUserParams{}
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error parsing body: %v", err))
	}

    user, err := u.userDB.GetUser(ctx, accountId)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error querying the user: %v", err))
    }

    if params.FirstName == "" {
        params.FirstName = user.FirstName
    }

    if params.LastName == "" {
        params.LastName = user.LastName
    }

    if params.Email == "" {
        params.Email = user.Email
    }

	if err := u.userDB.EditUser(ctx, params, accountId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error editing user: %v", err))
	}

	return c.JSON(fiber.Map{"message": "user edited"})
}
