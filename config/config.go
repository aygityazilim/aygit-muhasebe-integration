package config

import (
	"log"

	"github.com/joho/godotenv"
)

// InitConfig loads environment variables from .env file
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}
}
