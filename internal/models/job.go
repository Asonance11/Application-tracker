package models

import (
	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/types"
	"gorm.io/gorm"
)

func CreateJobStatusType(db *gorm.DB) error {
	return db.Exec(`
        DO $$ 
        BEGIN
            CREATE TYPE job_status AS ENUM ('applied', 'got_response', 'got_interview', 'failed_interview', 'rejected', 'got_offer');
        EXCEPTION
            WHEN duplicate_object THEN null;
        END $$;
    `).Error
}

func CreateJob(job *types.Job) error {
	return database.GetDB().Create(job).Error
}
