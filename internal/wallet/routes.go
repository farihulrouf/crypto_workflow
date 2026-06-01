package wallet

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {

	router.Post(
		"/wallets",
		handler.CreateWallet,
	)

	router.Get(
		"/wallets",
		handler.GetWallets,
	)
}
