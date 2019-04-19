package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// Env returns the .env variable
func Env(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}
