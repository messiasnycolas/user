package config

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	User string
	Pass string
	Name string
}

func init() {
	cfg = new(config)
}

// Load function get keys from the environment
func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		cfg.API = APIConfig{
			Port: "3000",
		}
	} else {
		cfg.API = APIConfig{
			Port: port,
		}
	}
	cfg.DB = DBConfig{
		User: os.Getenv("DB_USERNAME"),
		Pass: os.Getenv("DB_PASSWORD"),
		Name: os.Getenv("DB_NAME"),
	}
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPIConfig() APIConfig {
	return cfg.API
}
