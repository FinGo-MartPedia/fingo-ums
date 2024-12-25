package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	result := Env[key]
	if result == "" {
		return defaultValue
	}
	return result
}
