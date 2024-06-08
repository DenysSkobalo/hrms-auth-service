package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ServerConfig SConfig

type SConfig struct {
	ServerPort string
}

func LoadServerConfig() {
	err := godotenv.Load("configs/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ServerConfig = SConfig{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
