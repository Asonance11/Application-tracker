package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Asonance11/Application-tracker/internal/config"
	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/handlers"
	"github.com/Asonance11/Application-tracker/internal/middleware"
	"github.com/Asonance11/Application-tracker/internal/models"
	"github.com/Asonance11/Application-tracker/internal/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load env variables
	config.LoadEnvVariables()
	dsn := os.Getenv("DB_URL")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	//initialize db
	if err := database.InitDB(dsn); err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}
	db := database.GetDB()

	if err := models.CreateJobStatusType(db); err != nil {
		log.Fatalf("Failed to create job status type: %v", err)
	}

	// Check if tables exist
	if !db.Migrator().HasTable(&types.User{}) && !db.Migrator().HasTable(&types.Job{}) {
		if err := db.AutoMigrate(&types.User{}, &types.Job{}); err != nil {
			log.Fatalf("Failed to auto migrate: %v", err)
		}
	} else {
		log.Println("Tables already exist, skipping auto migration")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow your local frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	//Auth routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	//Main routes
	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		api.GET("/user", handlers.GetUser)
		api.GET("/jobs", handlers.ListJobApplications)
		api.POST("/jobs", handlers.CreateJobApplication)
		api.GET("/jobs/:id", handlers.GetJobByID)
		api.PUT("/jobs/:id", handlers.UpdateJobApplication)
		api.DELETE("/jobs/:id", handlers.DeleteJobApplication)
	}

	r.Run(fmt.Sprintf(":%s", port))
}
