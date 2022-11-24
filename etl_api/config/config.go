package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// API_PORT será sobrescrita com a variavel que esta declarada em .env
var API_PORT = ""

// LoadConfig seta a variavel que esta em .env na variavel API_PORT
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	API_PORT = os.Getenv("API_PORT")
}
