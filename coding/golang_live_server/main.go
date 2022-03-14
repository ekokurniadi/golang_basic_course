package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello guys")
	})
	router.Listen(":3000")
}
