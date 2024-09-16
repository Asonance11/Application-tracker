package handlers

import (
	"net/http"

	"github.com/Asonance11/Application-tracker/internal/models"
	"github.com/Asonance11/Application-tracker/internal/types"
	"github.com/gin-gonic/gin"
)

func CreateJobApplication(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	authenticatedUser, ok := user.(*types.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user data"})
		return
	}

	var newJob types.Job

	if err := c.ShouldBindJSON(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newJob.UserID = authenticatedUser.ID
	newJob.Status = types.StatusApplied

	if err := models.CreateJob(&newJob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job application"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job application created successfully"})
}

func ListJobApplications(c *gin.Context) {}

func UpdateJobApplication(c *gin.Context) {}

func DeleteJobApplication(c *gin.Context) {}
