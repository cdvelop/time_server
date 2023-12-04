package timeserver

import (
	"time"

	"github.com/cdvelop/strings"
)

// layout ej: 2006-01-02 15:04
func (t timeServer) ToDay(layout string) string {

	if layout == "" {
		layout = "2006-01-02"
	}

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
