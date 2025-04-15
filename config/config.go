package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")

		return
	}

	log.Println(".env file has been loaded")
}

func getString(env string, defaultVal string) string {
	val := strings.Trim(os.Getenv(env), " ")

	if val == "" {
		return defaultVal
	}

	return val
}

func getInt(env string, defaultVal int) int {
	val := strings.Trim(os.Getenv(env), " ")

	if val == "" {
		return defaultVal
	}

	i, err := strconv.Atoi(val)

	if err != nil {
		return defaultVal
	}

	return i
}