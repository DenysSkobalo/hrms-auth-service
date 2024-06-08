package routers

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/configs"
	"hrms-auth-service/internal/database/repositories"
	"hrms-auth-service/internal/transport/handlers"
)

func Setup(app *fiber.App) {
	db := configs.ConnectDB()
	userRepository := repositories.NewUserRepository(db)
	handler := handlers.NewHandler(db, *userRepository)

	app.Post("/accounts/signup", handler.SignUp)
	app.Post("/accounts/login", handler.Login)

	app.Get("/users/getAll", handler.GetAllUsers)
}
