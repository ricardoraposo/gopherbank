package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/db"
	"github.com/ricardoraposo/gopherbank/internal/handlers"
	"github.com/ricardoraposo/gopherbank/internal/middlewares"
)

func (s *FiberServer) getHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) RegisterRoutes() {
    // db stores
	accountDB := db.NewAccountStore(s.db)
	userDB := db.NewUserDB(s.db)
	transactionDB := db.NewTransactionStore(s.db, accountDB)

    // handlers
	accountsHandler := handlers.NewAccountHandler(accountDB, userDB)
	transactionHandler := handlers.NewTransactionHandler(transactionDB)
	authHandler := handlers.NewAuthHandler(accountDB)

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// accounts routes
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)

	// transactions routes
	api.Post("/transfer", transactionHandler.Transfer)
	api.Post("/withdraw", transactionHandler.Withdraw)

	// auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", authHandler.Authenticate)
	auth.Post("/new", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)

    // admin routes
	api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Post("/deposit", middlewares.IsAdmin, transactionHandler.Deposit)

	s.Get("/health", s.getHealth)
}
