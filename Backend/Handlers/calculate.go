package handlers

import (
	"calculateDiscount/requests"
	"calculateDiscount/services"

	"github.com/gofiber/fiber/v2"
)

func CalculateDiscount(c *fiber.Ctx) error {
	var req requests.Cart
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	total := services.ApplyDiscount(req)

	return c.JSON(fiber.Map{"total": total})
}
