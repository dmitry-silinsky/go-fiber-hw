package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ListenAddr string
}

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")

		return
	}

	log.Println(".env file has been loaded")
}

func NewAppConfig() *AppConfig {
	val, exists := os.LookupEnv("LISTEN_ADDR")
	if !exists {
		val = "localhost:3000"
	}

	return &AppConfig{
		ListenAddr: val,
	}
}