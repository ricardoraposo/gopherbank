package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/database"
)

type FiberServer struct {
	*fiber.App
	db *database.Store
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(),
		db:  database.New(),
	}
}
