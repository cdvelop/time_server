package timeserver

import "time"

//date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
func (TimeServer) WeekDayNumber(date_in string) (wd int, err string) {

	dateTime, e := time.Parse("2006-01-02", date_in)
	if e != nil {
		return 0, "WeekDayNumber error " + e.Error()
	}

	return int(dateTime.Weekday()), ""
}
