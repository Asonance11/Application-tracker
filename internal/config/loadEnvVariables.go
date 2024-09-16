package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env variables")
	}
}
