package timeserver

import (
	"strconv"
	"time"
)

func YearsToBirthDay(yearsTxt string) string {
	yearInt, err := strconv.Atoi(yearsTxt)
	if err != nil {
		return ""
	}
	now := time.Now()
	// fmt.Println("Today:", now)
	afterDate := now.AddDate(-yearInt, 0, 0)

	return afterDate.Format("2006-01-02")
}
