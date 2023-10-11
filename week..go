package timeserver

import "time"

//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
func (TimeServer) WeekDayNumber(date_in string) (int, error) {

	dateTime, err := time.Parse("2006-01-02", date_in)
	if err != nil {
		return 0, err
	}

	return int(dateTime.Weekday()), nil
}
