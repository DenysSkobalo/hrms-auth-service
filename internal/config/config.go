package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ServerConfig Config

type Config struct {
	ServerPort string
}

func LoadConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ServerConfig = Config{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
