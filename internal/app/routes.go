package app

import (
	"crypto-flow/internal/auth"
	"crypto-flow/internal/middleware"
	"crypto-flow/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,
	authHandler *auth.Handler,
	walletHandler *wallet.Handler,
) {

	api := app.Group("/api/v1")

	// Public Routes
	authGroup := api.Group("/auth")

	auth.RegisterRoutes(
		authGroup,
		authHandler,
	)

	// Protected Routes
	private := api.Group(
		"/private",
		middleware.JWT(),
	)

	private.Get(
		"/profile",
		authHandler.Profile,
	)

	wallet.RegisterRoutes(
		private,
		walletHandler,
	)
}
