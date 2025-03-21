package main

import (
	"calculateDiscount/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/", handlers.CalculateDiscount)
	app.Listen(":3000")
}
