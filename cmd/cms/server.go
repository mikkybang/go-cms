package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikkybang/go-cms/config"
	"github.com/mikkybang/go-cms/pkg/database"
)

func main() {
	app := fiber.New()

	config := config.New()

	database.Connect();

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + config.Port)
}
