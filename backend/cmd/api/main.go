package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ricardoraposo/gopherbank/internal/server"
)

func main() {
	app := server.New()
	app.RegisterRoutes()

    port := fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT"))
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
