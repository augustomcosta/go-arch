package main

import (
	"context"
	"github.com/augustomcosta/go-arch/config"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbcon := config.ConnectToDB()
	container := config.InitDIContainer(dbcon)
	app := config.NewApp(container)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}
}
