package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) SignUp(c *fiber.Ctx) error {
	return c.SendString("Signup Handler")
}