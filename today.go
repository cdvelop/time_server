package timeserver

import (
	"time"

	"github.com/cdvelop/timetools"
)

// layout ej: 2006-01-02 15:04
func (t TimeServer) ToDay(layout string) string {

	if layout == "" {
		layout = "2006-01-02"
	}

	err := timetools.CheckDateExists(t.CurrentTestDate)
	if err == "" {
		return t.CurrentTestDate
	}

	return time.Now().Format(layout)
}
