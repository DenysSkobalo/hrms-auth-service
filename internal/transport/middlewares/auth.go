package middlewares

import (
	"hrms-auth-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
			token := c.Get("Authorization")
			if token == "" {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
							"error": "Missing Authorization header",
					})
			}

			if !utils.IsValidToken(token) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
							"error": "Invalid token",
					})
			}

			return c.Next()
	}
}