package handlers

import (
	"net/http"
	"strconv"

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

func GetJobByID(c *gin.Context) {
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

	jobId := c.Param("id")

	var job types.Job

	if err := models.GetJobByID(jobId, &job); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job application not found"})
		return
	}

	if job.UserID != authenticatedUser.ID {
		c.JSON(
			http.StatusForbidden,
			gin.H{"error": "You do not have permission to update this job"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

func ListJobApplications(c *gin.Context) {
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

	// Parse pagination parameters (with defaults)
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	// Convert page and pageSize to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	// Calculate the offset for pagination
	offset := (pageInt - 1) * pageSizeInt

	result, err := models.GetJobApplicationsByUserID(authenticatedUser.ID, pageSizeInt, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": result, "page": pageInt, "pageSize": pageSizeInt})
}

func UpdateJobApplication(c *gin.Context) {
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

	jobId := c.Param("id")

	var job types.Job

	if err := models.GetJobByID(jobId, &job); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job application not found"})
		return
	}

	if job.UserID != authenticatedUser.ID {
		c.JSON(
			http.StatusForbidden,
			gin.H{"error": "You do not have permission to update this job"},
		)
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.UpdateJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job application updated succesfully"})

}

func DeleteJobApplication(c *gin.Context) {
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

	jobId := c.Param("id")

	var job types.Job

	if err := models.GetJobByID(jobId, &job); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job application not found"})
		return
	}

	if job.UserID != authenticatedUser.ID {
		c.JSON(
			http.StatusForbidden,
			gin.H{"error": "You do not have permission to delete this job"},
		)
		return
	}

	if err := models.DeleteJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job application deleted succesfully"})

}
