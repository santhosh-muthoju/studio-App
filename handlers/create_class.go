package handlers

import (
	"net/http"
	"studio-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateClass(c *gin.Context) {
	var classRequest models.ClassRequest
	if err := c.ShouldBindJSON(&classRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", classRequest.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", classRequest.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
		return
	}

	classID := uuid.New().String()

	class := models.Class{
		ID:        classID,
		ClassName: classRequest.ClassName,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  classRequest.Capacity,
	}

	models.AddClass(class)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Class created successfully",
		"data": gin.H{
			"class_id": class.ID,
			"class":    class.ClassName,
		},
	})
}
