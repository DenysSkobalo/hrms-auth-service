package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/database/repositories"
)

type IHandler interface {
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error

	GetAllUsers(c *fiber.Ctx) error
}

type Handler struct {
	DB             *sql.DB
	UserRepository repositories.UserRepository
}

func NewHandler(db *sql.DB, userRepository repositories.UserRepository) *Handler {
	return &Handler{
		DB:             db,
		UserRepository: userRepository,
	}
}
