package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.SendString("Login Handler")
}
