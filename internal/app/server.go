package app

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/configs"
	"hrms-auth-service/internal/transport/middlewares"
	"hrms-auth-service/internal/transport/routers"
	"log"
)

func Run() {
	configs.LoadServerConfig()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "HRMS v0.0.1",
	})

	app.Use(middlewares.CORSMiddleware())

	routers.Setup(app)

	log.Fatal(app.Listen(configs.ServerConfig.ServerPort))
}
