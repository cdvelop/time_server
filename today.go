package timeserver

import (
	"time"

	"github.com/cdvelop/strings"
)

// layout ej: 2006-01-02 15:04
func (t timeServer) DateToDay() string {
	return t.date("2006-01-02")
}

func (t timeServer) DateToDayHour() string {
	return t.date("2006-01-02 15:04:05")
}

func (t timeServer) date(layout string) string {

	if t.current_date != "" {
		var hour string
		if layout == "2006-01-02 15:04:05" && strings.Contains(t.current_date, ":") == 0 {
			ahora := time.Now()
			hour = " " + ahora.Format("15:04:05")
		}

		// 2006-01-02 15:04:05
		return t.current_date + hour
	}

	return time.Now().Format(layout)

}
