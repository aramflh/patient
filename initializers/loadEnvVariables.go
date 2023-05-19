// Package initializers provides additional functions for initializing the main function
package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	// err loads the environment variables defined in the .env file
	err := godotenv.Load()
	// nil is zero value for interface
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
