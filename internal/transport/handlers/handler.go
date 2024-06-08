package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error

	GetAllUsers(c *fiber.Ctx) error
}

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}
