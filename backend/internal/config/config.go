package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppVersion string
	Port       string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: os.Getenv("APP_VERSION"),
		Port:       os.Getenv("PORT"),
	}
}
