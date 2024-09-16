package date

import "time"

type IDateUtil interface {
	ConvertDateWithTimezoneToDateOnly(input string) (string, error)
	HasTimezone(dateStr string) bool
}

type DateUtil struct{}

func (d *DateUtil) ConvertDateWithTimezoneToDateOnly(input string) (string, error) {
	// Parse the input date string
	t, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		return "", err
	}

	// Create a new time object with the date part and set time to midnight in UTC
	dateOnly := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	// Format the new time object back to string
	return dateOnly.Format(time.RFC3339), nil
}

// HasTimezone checks if the given date string includes a timezone.
func (d *DateUtil) HasTimezone(dateStr string) bool {
	_, err := time.Parse(time.RFC3339, dateStr)
	return err == nil
}
