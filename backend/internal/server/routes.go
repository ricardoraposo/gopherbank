package server

import (
	"github.com/ricardoraposo/gopherbank/internal/handlers"
	"github.com/ricardoraposo/gopherbank/internal/middlewares"
)

func (s *FiberServer) RegisterRoutes() {
	// handlers
	accountsHandler := handlers.NewAccountHandler(s.db)
	transactionHandler := handlers.NewTransactionHandler(s.db)
	authHandler := handlers.NewAuthHandler(s.db)

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// accounts routes
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)
	// api.Get("accounts", accountsHandler.GetAllAccounts)

	// transactions routes
	api.Post("/transfer", transactionHandler.Transfer)
	api.Post("/withdraw", transactionHandler.Withdraw)
	// api.Post("/deposit", transactionHandler.Deposit)

	// auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", authHandler.Authenticate)
	auth.Post("/new", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)
	
	// admin routes
	api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Post("/deposit", middlewares.IsAdmin, transactionHandler.Deposit)
}
