package main

import (
	"log"

	"crypto-flow/internal/auth"
	"crypto-flow/internal/config"
	"crypto-flow/internal/database"
	"crypto-flow/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load .env
	config.LoadEnv()

	// Connect database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Auth dependencies
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// Fiber app
	app := fiber.New()

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// API v1
	api := app.Group("/api/v1")

	// Auth Routes
	authGroup := api.Group("/auth")
	auth.RegisterRoutes(authGroup, authHandler)

	// Protected Routes
	private := api.Group(
		"/private",
		middleware.JWT(),
	)

	private.Get(
		"/profile",
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "authenticated",
			})
		},
	)

	// Start Server
	port := config.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)

	log.Fatal(
		app.Listen(":" + port),
	)
}
