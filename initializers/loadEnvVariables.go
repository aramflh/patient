// Package initializers provides additional functions for initializing the main function
package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvVariables loads the variables in .env file
// They then  become accessible with : os.Getenv("ENV_VAR")
func LoadEnvVariables() {
	// err loads the environment variables defined in the .env file
	err := godotenv.Load()
	// nil is zero value for interface
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
