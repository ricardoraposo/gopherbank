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
	depositRequestHandler := handlers.NewDepositRequestHandler(s.db)
	notificationHandler := handlers.NewNotificationHandler(s.db)

	s.Use(logger.New())
	s.Use(cors.New())

	api := s.App.Group("/api")
	api.Use(middlewares.JWTAuthentication)

	// jwt routes
	api.Get("/jwt", middlewares.GetJWTAccount)

	// accounts routes
	api.Get("/accounts/:id", accountsHandler.GetAccountByNumber)
	api.Delete("/accounts/:id", accountsHandler.DeleteAccount)

	// user routes
	api.Get("/user/:id", userHandler.GetUser)
	api.Put("/user", userHandler.EditUser)

	// transactions routes
	api.Get("/transaction/:id", transactionHandler.GetAccountTransactions)
	api.Post("/transfer", middlewares.ValidateTransferParams, transactionHandler.Transfer)
	api.Post("/withdraw", middlewares.ValidateWithdrawParams, transactionHandler.Withdraw)

	// deposit request routes
	api.Post("/deposit-request", depositRequestHandler.CreateDepositRequest)

	// favorite routes
	api.Get("/favorite", favoritehandler.GetFavorites)
	api.Post("/favorite", favoritehandler.ToggleFavorite)

	// notification routes
	api.Get("/notification", notificationHandler.GetAccountNotifications)
	api.Delete("/notification/:id", notificationHandler.RemoveNotification)

	// auth routes
	auth := s.App.Group("/auth")
	auth.Post("/", middlewares.ValidateLoginParams, authHandler.Authenticate)
	auth.Post("/new", middlewares.ValidateNewAccountParams, accountsHandler.CreateAccount)
	auth.Post("/new/default", middlewares.ValidateNewAccountParamsNoS3, accountsHandler.CreateAccountNoS3)
	auth.Patch("/recover", middlewares.ValidateNewPasswordParams, accountsHandler.RecoverAccountPassword)

	// admin routes
	// api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Get("/accounts", middlewares.IsAdmin, accountsHandler.GetAllAccounts)
	api.Get("/transaction", middlewares.IsAdmin, transactionHandler.GetAllTransactions)
	api.Get("/deposit-request", middlewares.IsAdmin, depositRequestHandler.GetAllDepositRequests)
	api.Post("/deposit-request/approve/:id", middlewares.IsAdmin, depositRequestHandler.ApproveRequest)
	api.Patch("/deposit-request/reject/:id", middlewares.IsAdmin, depositRequestHandler.RejectRequest)
	api.Post("/deposit", middlewares.IsAdmin, transactionHandler.Deposit)
}
