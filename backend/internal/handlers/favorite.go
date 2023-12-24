package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/models"
)

type FavoriteHandler struct {
	db db.FavoriteDB
}

func NewFavoriteHandler(client *db.DB) *FavoriteHandler {
	favoriteDB := db.NewFavoriteDB(client)
	return &FavoriteHandler{
		db: favoriteDB,
	}
}

func (f *FavoriteHandler) AddToFavorite(c *fiber.Ctx) error {
	p := models.NewFavoriteParams{}
	if err := c.BodyParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		fmt.Println("could not get the claims")
		return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
	}

	if claims["number"] != p.AccountID {
		return fiber.NewError(fiber.StatusUnauthorized, "Not enough credentials")
	}

	if err := f.db.CreateFavorite(c.Context(), p); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("%v and %v are now besto friendos", p.AccountID, p.FavoritedID),
	})
}
