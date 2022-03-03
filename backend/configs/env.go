package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, defaulting to exported variables.")
	}
}
