package handlers

import (
	"calculateDiscount/requests"
	"calculateDiscount/services"

	"github.com/gofiber/fiber/v2"
)

func CalculateDiscount(c *fiber.Ctx) error {
	var req requests.Cart

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid input",
			"message": err.Error(),
		})
	}

	total, err := services.ApplyDiscount(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid input",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Net Price": total,
	})
}
