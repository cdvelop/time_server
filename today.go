package timeserver

import (
	"time"

	"github.com/cdvelop/timetools"
)

func (t timeServer) DateToDay() string {

	date, _ := timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date)

	return date
}

func (t timeServer) DateToDayHour(seconds ...bool) (date, hour string) {

	return timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date, seconds...)
}
