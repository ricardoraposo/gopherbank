package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ricardoraposo/gopherbank/internal/server"
)

func main() {
	app := server.New()
	app.RegisterRoutes()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	go func() {
		if err := app.Listen(port); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		panic(err)
	}

	fmt.Println("\nServer gracefully stopped")
}
