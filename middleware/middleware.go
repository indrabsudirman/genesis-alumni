package middleware

import (
	"genesis-alumni/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("key-access")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	//Send to AlumnusHandlerGetAll
	ctx.Locals("userInfo", claims)
	// ctx.Locals("role", claims["role"])

	return ctx.Next()

}
