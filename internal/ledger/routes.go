package ledger

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {

	router.Post(
		"/ledgers",
		handler.CreateLedger,
	)

	router.Get(
		"/ledgers",
		handler.GetLedgers,
	)
}
