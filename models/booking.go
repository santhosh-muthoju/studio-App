package models

import "time"

type BookingRequest struct {
	ClassID    string `json:"class_id"`
	MemberName string `json:"member_name"`
	Date       string `json:"date"`
}

type Booking struct {
	ClassID    string
	ClassName  string
	MemberName string
	Date       time.Time
}

var bookings []Booking

func AddBooking(booking Booking) {
	bookings = append(bookings, booking)
}

func CheckCapacity(classID string, date time.Time) bool {
	var count int
	for _, booking := range bookings {
		if booking.ClassID == classID && booking.Date.Equal(date) {
			count++
		}
	}

	class, _ := FindClassByID(classID)
	return count < class.Capacity
}
