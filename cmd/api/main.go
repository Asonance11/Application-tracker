package main

import (
	"os"

	"github.com/Asonance11/Application-tracker/internal/config"
	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/handlers"
	"github.com/Asonance11/Application-tracker/internal/middleware"
	"github.com/Asonance11/Application-tracker/internal/models"
	"github.com/Asonance11/Application-tracker/internal/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load env variables
	config.LoadEnvVariables()
	dsn := os.Getenv("DB_URL")

	//initialize db
	if err := database.InitDB(dsn); err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}

	db := database.GetDB()

	if err := models.CreateJobStatusType(db); err != nil {
		log.Fatalf("Failed to create job status type: %v", err)
	}

	if !db.Migrator().HasTable(&types.User{}) {
		if err := db.AutoMigrate(&types.User{}, &types.Job{}); err != nil {
			log.Fatalf("Failed to auto migrate: %v", err)
		}
	} else {
		log.Info("Users table already exists, skipping migration")
	}

	r := gin.Default()

	//Auth routes

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	//Main routes
	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		api.POST("/jobs", handlers.CreateJobApplication)
	}

	r.Run()
}
