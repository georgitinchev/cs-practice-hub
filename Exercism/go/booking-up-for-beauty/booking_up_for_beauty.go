package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	formats := []string{
		"1/2/2006 15:04:05",
		"1/2/2006 15:04",
		"January 2, 2006 15:04:05",
		"January 2, 2006 15:04",
		"2 Jan 2006 15:04:05",
		"2 Jan 2006 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
	}
	for _, format := range formats {
		if t, err := time.Parse(format, date); err == nil {
			return t
		}
	}
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	scheduledTime := Schedule(date)
	currentTime := time.Now()
	return scheduledTime.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	formats := []string{
		"1/2/2006 15:04:05",
		"1/2/2006 15:04",
		"January 2, 2006 15:04:05",
		"January 2, 2006 15:04",
		"2 Jan 2006 15:04:05",
		"2 Jan 2006 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"Monday, January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04",
		"Monday, January 2, 2006 15:04",
		"Monday, March 8, 1974 12:02:02",
	}
	for _, format := range formats {
		if t, err := time.Parse(format, date); err == nil {
			hour := t.Hour()
			return hour >= 12 && hour < 18
		}
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	scheduledTime := Schedule(date)
	if scheduledTime.IsZero() {
		return ""
	}
	return "You have an appointment on " + scheduledTime.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
