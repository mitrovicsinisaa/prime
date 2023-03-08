package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mitrovicsinisaa/prime/primes"
)

func main() {
	address := flag.String("address", ":3000", "listen address")
	flag.Parse()

	app := fiber.New()
	app.Post("/prime", primes.CheckPrimes)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := app.Listen(*address); err != nil {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	<-quit
	log.Println("Server is shutting down...")

	if err := app.Server().ShutdownWithContext(ctx); err != nil {
		log.Fatalf("error shutting down server: %v\n", err)
	}

	log.Println("Server shutdown completed.")
}
