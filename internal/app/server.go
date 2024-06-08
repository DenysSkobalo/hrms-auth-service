package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hrms-auth-service/internal/config"
	"hrms-auth-service/internal/transport/routers"
	"log"
)

func Run() {
	db, err := config.ConnectDB()
	if err != nil {
		err = fmt.Errorf("failed to connect database: %w", err)
	}
	defer db.Close()

	config.LoadServerConfig()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "HRMS v0.0.1",
	})

	routers.Setup(app, db)

	log.Fatal(app.Listen(config.ServerConfig.ServerPort))
}
