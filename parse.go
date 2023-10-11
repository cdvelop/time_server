package timeserver

import "time"

// formato fecha "2006-01-02"
func ParseDate(date_in string) (time.Time, error) {
	dateTime, err := time.Parse("2006-01-02", date_in)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}
