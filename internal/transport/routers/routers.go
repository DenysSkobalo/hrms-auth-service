package routers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/transport/handlers"
)

func Setup(app *fiber.App, db *sql.DB) {
	handler := handlers.NewHandler(db)

	app.Post("/accounts/signup", handler.SignUp)
	app.Post("/accounts/login", handler.Login)
}
