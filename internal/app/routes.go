package app

import (
	"crypto-flow/internal/auth"
	"crypto-flow/internal/ledger"
	"crypto-flow/internal/middleware"
	"crypto-flow/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,
	authHandler *auth.Handler,
	walletHandler *wallet.Handler,
	ledgerHandler *ledger.Handler,
) {

	api := app.Group("/api/v1")

	authGroup := api.Group("/auth")
	auth.RegisterRoutes(
		authGroup,
		authHandler,
	)

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

	ledger.RegisterRoutes(
		private,
		ledgerHandler,
	)
}
