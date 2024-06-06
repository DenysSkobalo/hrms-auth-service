package routers

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/transport/handlers"
)

func Setup(app *fiber.App) {
	handler := handlers.NewHandler()

	app.Post("/accounts/signup", handler.SignUp)
	app.Post("/accounts/login", handler.Login)
}
