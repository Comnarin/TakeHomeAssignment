package main

import (
	"calculateDiscount/handlers"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()
	app.Get("/",handlers.CalculateDiscount)
	app.Listen(":3000")
}
