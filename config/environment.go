package config

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

// LoadEnv extract keys from the .env file
func LoadEnv(keys ...string) (values []string) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for _, key := range keys {
		value := os.Getenv(key)
		values = append(values, value)
	}

	return
}
