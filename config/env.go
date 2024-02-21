package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Retrieves environment variable with a specified key
func GoEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
