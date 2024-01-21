package timeserver

import (
	"time"

	"github.com/cdvelop/timetools"
)

func (t timeServer) DateToDay(df timetools.DateFormatAdapter) string {
	date, _ := timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date, df)
	return date
}

func (t timeServer) DateToDayHour(df timetools.DateFormatAdapter) (date, hour string) {
	return timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date, df)
}
