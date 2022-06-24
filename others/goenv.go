package others

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := env.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
