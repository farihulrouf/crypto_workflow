package main

import (
	"log"

	"crypto-flow/internal/app"
	"crypto-flow/internal/config"
	"crypto-flow/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Load Environment
	config.LoadEnv()

	// Database Connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// Dependency Container
	container := app.NewContainer(db)

	// Fiber
	appFiber := fiber.New()

	// Middleware
	appFiber.Use(cors.New())

	// Health Check
	appFiber.Get(
		"/health",
		func(c *fiber.Ctx) error {
			return c.JSON(
				fiber.Map{
					"status": "ok",
				},
			)
		},
	)

	// Routes
	app.RegisterRoutes(
		appFiber,
		container.AuthHandler,
		container.WalletHandler,
		container.LedgerHandler,
	)

	// Port
	port := config.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf(
		"Server running on :%s",
		port,
	)

	log.Fatal(
		appFiber.Listen(
			":" + port,
		),
	)
}
