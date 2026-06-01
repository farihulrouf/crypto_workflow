package auth

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(
	c *fiber.Ctx,
) error {

	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid payload",
			},
		)
	}

	err := h.service.Register(req)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "register success",
		},
	)
}

func (h *Handler) Login(
	c *fiber.Ctx,
) error {

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid payload",
			},
		)
	}

	token, err := h.service.Login(req)

	if err != nil {
		return c.Status(401).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"token": token,
		},
	)
}

func (h *Handler) Profile(
	c *fiber.Ctx,
) error {

	return c.JSON(
		fiber.Map{
			"user_id": c.Locals("user_id"),
		},
	)
}
