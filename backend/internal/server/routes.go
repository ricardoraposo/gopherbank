package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/database"
	"github.com/ricardoraposo/gopherbank/internal/handlers"
	"github.com/ricardoraposo/gopherbank/internal/middlewares"
)

func (s *FiberServer) getHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) RegisterRoutes() {
	accountStore := database.NewAccountStore(s.db)
	accountsHandler := handlers.NewAccountHandler(accountStore)
	authHandler := handlers.NewAuthHandler(accountStore)

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// accounts routes
	api.Get("/accounts", accountsHandler.GetAllAccounts)
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Post("/accounts", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)

	//auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", authHandler.Authenticate)

	s.Get("/health", s.getHealth)
}
