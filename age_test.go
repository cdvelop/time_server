package timeserver_test

import (
	"testing"
	"time"

	"github.com/cdvelop/timeserver"
)

type AgeTestCandidate struct {
	BirthDate    time.Time
	CheckingTime time.Time
	ExpectedAge  int
}

var AgeTestCandidates = []AgeTestCandidate{
	{time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 10},
	{time.Date(2001, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 8},
	{time.Date(2004, 6, 18, 0, 0, 0, 0, time.UTC), time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 0},
	{time.Date(1981, 3, 21, 0, 0, 0, 0, time.UTC), time.Date(2021, 3, 20, 0, 0, 0, 0, time.UTC), 39},
	{time.Date(1981, 3, 21, 0, 0, 0, 0, time.UTC), time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC), 41},
}

func TestAgeAt(t *testing.T) {
	for _, candidate := range AgeTestCandidates {
		gotAge := timeserver.AgeAt(candidate.BirthDate, candidate.CheckingTime)
		if gotAge != candidate.ExpectedAge {
			t.Error(
				"For", candidate.BirthDate,
				"Expected", candidate.ExpectedAge,
				"Got", gotAge,
			)
		}
	}
}
