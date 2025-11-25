package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func GetEnv(key string, defaulValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaulValue
	}

	return value
}
