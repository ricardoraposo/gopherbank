package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gopherbank/internal/db"
)

type FiberServer struct {
	*fiber.App
	db *db.DB
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(),
		db:  db.New(),
	}
}
