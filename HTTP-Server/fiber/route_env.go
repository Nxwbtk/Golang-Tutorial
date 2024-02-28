package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getEnv(c *fiber.Ctx) error {
	// this is for env on local
	// if value, exists := os.LookupEnv("SECRET") ; exists {
	// 	return c.JSON(fiber.Map{
	// 		"SECRET": value,
	// 	})
	// }
	// return c.JSON(fiber.Map{
	// 	"SECRET": "Not found",
	// })
	return c.JSON(fiber.Map{
		"SECRET_ENV_FILE": os.Getenv("SECRET_ENV"),
	})
}