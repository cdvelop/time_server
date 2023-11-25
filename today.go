package timeserver

import (
	"time"
)

// layout ej: 2006-01-02 15:04
func (t timeServer) ToDay(layout string) string {

	if layout == "" {
		layout = "2006-01-02"
	}

	if t.current_date != "" {
		return t.current_date
	}

	return time.Now().Format(layout)
}
