package main

import (
	"log"

	"crypto-flow/internal/app"
	"crypto-flow/internal/auth"
	"crypto-flow/internal/config"
	"crypto-flow/internal/database"
	"crypto-flow/internal/wallet"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.LoadEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	walletRepo := wallet.NewRepository(db)
	walletService := wallet.NewService(walletRepo)
	walletHandler := wallet.NewHandler(walletService)

	appFiber := fiber.New()

	appFiber.Use(cors.New())

	appFiber.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.RegisterRoutes(
		appFiber,
		authHandler,
		walletHandler,
	)

	port := config.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)

	log.Fatal(
		appFiber.Listen(":" + port),
	)
}
