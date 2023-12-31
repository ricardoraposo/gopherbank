package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	accountID := c.Params("id")
	user, err := u.userDB.GetUser(c.Context(), accountID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error querying the user: %v", err))
	}

	return c.JSON(fiber.Map{"user": user})
}

func (u *UserHandler) EditUser(c *fiber.Ctx) error {
	ctx := c.Context()
	params := models.EditUserParams{}
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error parsing body: %v", err))
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("could not get claims"))
	}

    id, ok := claims["number"].(string)
    fmt.Println(id)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid token"))
	}

	user, err := u.userDB.GetUser(ctx, id)
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

	if params.PictureURL == "" {
		params.PictureURL = user.PictureURL
	}

	if err := u.userDB.EditUser(ctx, params, id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("error editing user: %v", err))
	}

	return c.JSON(fiber.Map{"message": "user edited"})
}
