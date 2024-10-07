package handlers

import (
	"net/http"
	"studio-app/models"

	"github.com/gin-gonic/gin"
)

type BookResponse struct {
	BookingId  string `json:"booking_id"`
	ClassName  string `json:"class"`
	BookedDate string `json:"date"`
	Name       string `json:"member_name"`
}

func ListOfBookings(c *gin.Context) {
	var bookingData []BookResponse

	for _, booking := range models.Bookings {
		bookingData = append(bookingData, BookResponse{
			BookingId:  booking.BookingId,
			ClassName:  booking.ClassName,
			BookedDate: booking.Date.Format("2006-01-02T15:04:05Z"),
			Name:       booking.MemberName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Bookings fetched",
		"data":    bookingData,
	})
}
