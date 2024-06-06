package handlers

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
