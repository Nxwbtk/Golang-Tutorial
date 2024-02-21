package main

import (
	"github.com/gofiber/fiber/v2"
)

// https://youtu.be/nfSbFk8y4D0?si=nrHMvbt6BXB8rorp&t=773
func main() {
	app := fiber.New()
	app.Get("/hello/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
	})

	app.Listen(":8080")
}