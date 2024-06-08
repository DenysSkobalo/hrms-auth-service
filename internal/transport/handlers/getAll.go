package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepository.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
