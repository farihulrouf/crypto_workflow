package wallet

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

func (h *Handler) CreateWallet(
	c *fiber.Ctx,
) error {

	var req CreateWalletRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid payload",
			},
		)
	}

	userID := c.Locals("user_id").(string)

	err := h.service.CreateWallet(
		userID,
		req.Asset,
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "wallet created",
		},
	)
}

func (h *Handler) GetWallets(
	c *fiber.Ctx,
) error {

	userID := c.Locals("user_id").(string)

	wallets, err := h.service.GetWallets(
		userID,
	)

	if err != nil {
		return err
	}

	return c.JSON(wallets)
}
