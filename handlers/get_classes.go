package handlers

import (
	"net/http"
	"studio-app/models"

	"github.com/gin-gonic/gin"
)

type ClassResponse struct {
	ID        string `json:"id"`
	ClassName string `json:"class"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func ListOfClasses(c *gin.Context) {

	var responseData []ClassResponse

	for _, class := range models.ClassesList {
		responseData = append(responseData, ClassResponse{
			ID:        class.ID,
			ClassName: class.ClassName,
			StartDate: class.StartDate.Format("2006-01-02T15:04:05Z"),
			EndDate:   class.EndDate.Format("2006-01-02T15:04:05Z"),
			Capacity:  class.Capacity,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Classes fetched",
		"data":    responseData,
	})
}
