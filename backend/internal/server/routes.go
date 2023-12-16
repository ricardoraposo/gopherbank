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

    // accounts routes
	api.Get("/accounts", accountsHandler.GetAllAccounts)
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Post("/accounts", accountsHandler.CreateAccount)
    api.Delete("/accounts/:id", accountsHandler.DeleteAccount)

	s.Get("/health", s.getHealth)
}
