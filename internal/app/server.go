package app

import (
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/config"
	"hrms-auth-service/internal/transport/routers"
	"log"
)

func Run() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "HRMS v0.0.1",
	})

	routers.Setup(app)

	log.Fatal(app.Listen(config.ServerConfig.ServerPort))
}
