package server

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	s.Use(logger.New())
	s.Use(cors.New())

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// accounts routes
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)
	api.Patch("/accounts/", accountsHandler.RecoverAccountPassword)

	// user routes
	api.Put("/user/:id", userHandler.EditUser)

	// transactions routes
	api.Get("/transaction/:id", transactionHandler.GetAccountTransactions)
	api.Post("/transfer", middlewares.ValidateTransferParams, transactionHandler.Transfer)
	api.Post("/withdraw", middlewares.ValidateWithdrawParams, transactionHandler.Withdraw)

	// favorite routes
	api.Post("/favorite", favoritehandler.AddToFavorite)

	// auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", middlewares.ValidateLoginParams, authHandler.Authenticate)
	auth.Post("/new", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)

	// admin routes
	// api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Get("/transaction", middlewares.IsAdmin, transactionHandler.GetAllTransactions)
	api.Post("/deposit", middlewares.IsAdmin, transactionHandler.Deposit)
}
