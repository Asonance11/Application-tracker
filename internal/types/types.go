package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Jobs     []Job `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Job struct {
	gorm.Model
	Role           string
	CompanyName    string `gorm:"column:company_name"`
	ExpectedSalary uint
	Status         JobStatus `gorm:"type:job_status"`
	UserID         uint
}

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
