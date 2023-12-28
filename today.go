package timeserver

import (
	"time"

	"github.com/cdvelop/model"
	"github.com/cdvelop/timetools"
)

func (t timeServer) DateToDay(df *model.DateFormat) string {
	date, _ := timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date, df)
	return date
}

func (t timeServer) DateToDayHour(df *model.DateFormat) (date, hour string) {
	return timetools.DateToDayHour(time.Now().Format("2006-01-02 15:04:05"), t.fake_date, df)
}
