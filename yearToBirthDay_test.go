package timeserver_test

import (
	"testing"
)

var (
	dataAge = map[string]struct {
		inputAge    string
		outBirthday string
	}{
		"años":      {"29", "1991-09-15"},
		"años 40":   {"40", "1981-09-15"},
		" no faño ": {"", "1981-09-15"},
	}
)

func Test_YearsToBirthDay(t *testing.T) {
	for prueba, data := range dataAge {
		t.Run((prueba + data.inputAge), func(t *testing.T) {
			// birthday := timeserver.YearsToBirthDay(data.inputAge)

			// if birthday != data.outBirthday {
			// 	t.Fatalf("error se esperaba: [%v]\n-pero se obtuvo: [%v]", data.outBirthday, birthday)
			// }

		})
	}
}
