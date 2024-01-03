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

func (f *FavoriteHandler) ToggleFavorite(c *fiber.Ctx) error {
	p := models.NewFavoriteParams{}
	if err := c.BodyParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	claims, ok := c.Context().Value("claims").(jwt.MapClaims)
	if !ok {
		fmt.Println("could not get the claims")
		return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
	}

    p.AccountID, ok = claims["number"].(string)
    if !ok {
        fmt.Println("could not get the account number")
        return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
    }

	if err := f.db.ToggleFavorite(c.Context(), p); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("%v and %v are now besto friendos", p.AccountID, p.FavoritedID),
	})
}

func (f *FavoriteHandler) GetFavorites(c *fiber.Ctx) error {
    claims, ok := c.Context().Value("claims").(jwt.MapClaims)
    if !ok {
        fmt.Println("could not get the claims")
        return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
    }

    accountNumber, ok := claims["number"].(string)
    if !ok {
        fmt.Println("could not get the account number")
        return fiber.NewError(fiber.StatusInternalServerError, "authentication error")
    }

    favorites, err := f.db.GetFavoritedsByAccount(c.Context(), accountNumber)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    return c.JSON(fiber.Map{
        "favorites": favorites,
    })
}
