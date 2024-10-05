package handlers

import (
	"net/http"
	"studio-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func BookClass(c *gin.Context) {
	var bookingRequest models.BookingRequest
	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class, exists := models.FindClassByID(bookingRequest.ClassID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	bookingDate, err := time.Parse("2006-01-02", bookingRequest.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	booking := models.Booking{
		ClassID:    bookingRequest.ClassID,
		ClassName:  class.ClassName,
		MemberName: bookingRequest.MemberName,
		Date:       bookingDate,
	}

	if bookingDate.Before(class.StartDate) || bookingDate.After(class.EndDate) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Booking date is outside the class duration",
		})
		return
	}

	if !models.CheckCapacity(booking.ClassID, bookingDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class is fully booked for the selected date"})
		return
	}

	models.AddBooking(booking)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Booking successful",
		"data": gin.H{
			"class":       class.ClassName,
			"date":        bookingRequest.Date,
			"member_name": bookingRequest.MemberName,
		},
	})
}
