package handlers

import (
	"calculateDiscount/Requests"
	"calculateDiscount/Services"

	"github.com/gofiber/fiber/v2"
)

func CalculateDiscount2(c *fiber.Ctx) error {
	var req requests.Cart

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	total := services.ApplyDiscount(req)

	return c.JSON(fiber.Map{"total": total})
}
func CalculateDiscount(c *fiber.Ctx) error {
	var req requests.Cart

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	total := services.ApplyDiscount(req)

	return c.JSON(fiber.Map{"total": total})
}
