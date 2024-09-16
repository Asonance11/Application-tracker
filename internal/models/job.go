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
	if err := database.GetDB().Create(job).Error; err != nil {
		return err
	}
	return nil
}

func GetJobApplicationsByUserID(userID uint, limit int, offset int) ([]types.Job, error) {
	var jobs []types.Job

	if err := database.GetDB().Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}
