package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/constants/errors"
	"hrms-auth-service/internal/models"
	"hrms-auth-service/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errors.InvalidRequests})
	}

	exists, err := h.UserRepository.IsEmailExists(newUser.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.InternalServerError})
	}
	if exists {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": errors.ErrEmailExists})
	}

	newUser.Password = utils.HashPassword(newUser.Password)
	newUser.CreatedAt = time.Now()

	err = h.UserRepository.CreateUser(&newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errors.FailedToCreateUser})
	}

	return c.JSON(newUser)
}
