package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/database"
	"github.com/ricardoraposo/gopherbank/internal/handlers"
)

func (s *FiberServer) getHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) RegisterRoutes() {
	accountStore := database.NewAccountStore(s.db)
	accountsHandler := handlers.NewAccountHandler(accountStore)

	api := s.App.Group("/api")

	api.Get("/accounts", accountsHandler.GetAllAccounts)
	api.Post("/accounts", accountsHandler.CreateAccount)

	s.Get("/health", s.getHealth)
}
