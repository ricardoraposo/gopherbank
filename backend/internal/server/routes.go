package server

import (
	"github.com/ricardoraposo/gopherbank/internal/handlers"
	"github.com/ricardoraposo/gopherbank/internal/middlewares"
)

func (s *FiberServer) RegisterRoutes() {
	// handlers
	accountsHandler := handlers.NewAccountHandler(s.db)
	userHandler := handlers.NewUserHandler(s.db)
	transactionHandler := handlers.NewTransactionHandler(s.db)
	authHandler := handlers.NewAuthHandler(s.db)
	favoritehandler := handlers.NewFavoriteHandler(s.db)

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// accounts routes
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)

	// user routes
	api.Put("/user/:id", userHandler.EditUser)

	// transactions routes
	api.Get("/transaction/:id", transactionHandler.GetAccountTransactions)
	api.Post("/transfer", transactionHandler.Transfer)
	api.Post("/withdraw", transactionHandler.Withdraw)

	// favorite routes
	api.Post("/favorite", favoritehandler.AddToFavorite)

	// auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", authHandler.Authenticate)
	auth.Post("/new", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)

	// admin routes
	api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Get("/transaction", middlewares.IsAdmin, transactionHandler.GetAllTransactions)
	api.Post("/deposit", middlewares.IsAdmin, transactionHandler.Deposit)
}
