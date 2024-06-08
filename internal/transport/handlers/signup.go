package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/models"
	"hrms-auth-service/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if exists, err := h.UserRepository.IsEmailExists(newUser.Email); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	} else if exists {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
	}

	newUser.Password = utils.HashPassword(newUser.Password)
	newUser.CreatedAt = time.Now()

	err := h.UserRepository.CreateUser(&newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(newUser)
}
