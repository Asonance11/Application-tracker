package models

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

// JobStatus represents the status of a job application
type JobStatus string

// Define constants for each possible status
const (
	StatusApplied         JobStatus = "applied"
	StatusGotResponse     JobStatus = "got_response"
	StatusGotInterview    JobStatus = "got_interview"
	StatusFailedInterview JobStatus = "failed_interview"
	StatusRejected        JobStatus = "rejected"
	StatusGotOffer        JobStatus = "got_offer"
)

// Scan implements the sql.Scanner interface
func (js *JobStatus) Scan(value interface{}) error {
	*js = JobStatus(value.(string))
	return nil
}

// Value implements the driver.Valuer interface
func (js JobStatus) Value() (driver.Value, error) {
	return string(js), nil
}

// Job represents a job application
type Job struct {
	gorm.Model
	Role           string
	CompanyName    string `gorm:"column:company_name"`
	ExpectedSalary uint
	Status         JobStatus `gorm:"type:job_status"`
}

// CreateJobStatusType creates the job_status enum type in the database
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
