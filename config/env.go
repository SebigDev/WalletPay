package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
