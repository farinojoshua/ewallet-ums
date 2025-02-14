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
		log.Fatal("failed to load env", err)
	}
}

func GetEnv(key string, value string) string {
	// value used as a default
	result := Env[key]
	if result == "" {
		result = value
	}
	return result
}
