package main

import (
	"github.com/Asonance11/Application-tracker/internal/config"
	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load env variables
	config.LoadEnvVariables()

	//initialize db
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}

	r := gin.Default()

	//Auth routes

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.Run()
}
