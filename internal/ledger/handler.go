package ledger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

func (h *Handler) CreateLedger(
	c *fiber.Ctx,
) error {

	var req CreateLedgerRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid payload",
			},
		)
	}

	userIDStr := c.Locals("user_id").(string)

	userID, err := uuid.Parse(
		userIDStr,
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid user id",
			},
		)
	}

	err = h.service.CreateLedger(
		userID,
		req,
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
			"message": "ledger created",
		},
	)
}

func (h *Handler) GetLedgers(
	c *fiber.Ctx,
) error {

	userIDStr := c.Locals("user_id").(string)

	userID, err := uuid.Parse(
		userIDStr,
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid user id",
			},
		)
	}

	ledgers, err := h.service.GetLedgers(
		userID,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		ledgers,
	)
}
