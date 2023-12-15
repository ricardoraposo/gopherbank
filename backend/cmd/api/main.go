package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
