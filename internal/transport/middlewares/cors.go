package middlewares

import "github.com/gofiber/fiber/v2"

func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	}
}
