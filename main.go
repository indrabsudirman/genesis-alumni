package main

import (
	"genesis-alumni/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initial Database
	database.DatabaseInit()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "Genesis",
		})
	})
	app.Listen(":8080")
}
