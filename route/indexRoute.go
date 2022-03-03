package route

import "github.com/gofiber/fiber/v2"

func RouteInit(route *fiber.App) {

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"كَيْفَ حَالُكُمْ": "Genesis",
		})
	})
}