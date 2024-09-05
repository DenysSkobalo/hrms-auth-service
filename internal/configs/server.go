package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ServerConfig SConfig
var JWTSecret string

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

	JWTSecret = os.Getenv("JWT_SECRET")
    if JWTSecret == "" {
        log.Fatal("JWT_SECRET environment variable is required but not set")
    }
}
