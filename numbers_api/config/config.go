package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// APIPort receives the value from .env
var APIPort = ""

// Load configurations for API
func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APIPort = os.Getenv("API_PORT")
}
