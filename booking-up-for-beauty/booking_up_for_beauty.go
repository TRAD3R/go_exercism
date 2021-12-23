package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	appointment, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return appointment
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	passed, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return passed.Unix() < time.Now().Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	appointment, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appointmentTime := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s.", appointmentTime.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	dateInCurYear := fmt.Sprintf("%d-09-15", time.Now().Year())
	aniversary, _ := time.Parse("2006-01-02", dateInCurYear)

	return aniversary
}
