package database

import (
	"os"

	"github.com/Asonance11/Application-tracker/internal/config"
	"github.com/Asonance11/Application-tracker/internal/models"
	// log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	config.LoadEnvVariables()

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := models.CreateJobStatusType(db); err != nil {
		return err
	}

	// Auto migrate your models

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Job{}); err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
