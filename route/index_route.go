package route

import (
	"genesis-alumni/handler"
	"genesis-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"كَيْفَ حَالُكُمْ": "Genesis",
		})
	})

	route.Post("/login", handler.LoginHandler)
	route.Post("/alumnus", handler.AlumnusHandlerCreate)
	route.Get("/alumnus", middleware.Auth, handler.AlumnusHandlerGetAll)
}
