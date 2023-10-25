package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikkybang/go-cms/config"
)

func main() {
	app := fiber.New()

	config := config.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + config.Port)
}
