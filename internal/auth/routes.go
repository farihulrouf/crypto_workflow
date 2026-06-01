package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {
	router.Post(
		"/register",
		handler.Register,
	)

	router.Post(
		"/login",
		handler.Login,
	)
}

func RegisterProtectedRoutes(
	router fiber.Router,
	handler *Handler,
) {
	router.Get(
		"/profile",
		handler.Profile,
	)
}
