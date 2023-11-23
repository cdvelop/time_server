package timeserver_test

import (
	"testing"

	"github.com/cdvelop/timeserver"
	"github.com/cdvelop/timetools"
)

var (
	dataDateShort = map[string]struct {
		input          string
		dayTxtExpected string //Lun,Mar..Dom
		dayNumExpected string //31
		monthExpected  string //Ene,Feb..Dic
		yearExpected   string //2012
	}{
		"fecha 1 Sabado ok":     {"2021-09-18", "Sab", "18", "Sep", "21"},
		"fecha 2 domingo ok":    {"2021-07-25", "Dom", "25", "Jul", "21"},
		"fecha 3 Martes ok":     {"2019-03-04", "Lun", "4", "Mar", "19"},
		"fecha 4 Martes ok":     {"2024-03-01", "Vie", "1", "Mar", "24"},
		"fecha 5 Miercoles ok":  {"2025-01-01", "Mie", "1", "Ene", "25"},
		"fecha 6 format error":  {"2025-1-01", "StringDateToShort formato de fecha ingresado incorrecto ej: 2006-01-02", "", "", ""},
		"fecha 7 no data error": {"", "StringDateToShort formato de fecha ingresado incorrecto ej: 2006-01-02", "", "", ""},
	}
)

func Test_ShortDate(t *testing.T) {

	const text_error = "*error\n-se esperaba: [%v]\n-pero se obtuvo: [%v]\n-en la entrada: [%v]"

	for prueba, data := range dataDateShort {
		t.Run((prueba), func(t *testing.T) {

			timeHandler := timeserver.TimeServer{}

			txtD, D, txtM, Y, err := timetools.StringDateToShort(data.input, timeHandler)
			if err != "" {
				txtD = err
			}

			// and error handler
			if txtD != data.dayTxtExpected {
				t.Fatalf(text_error, data.dayTxtExpected, txtD, data.input)
			}

			if D != data.dayNumExpected {
				t.Fatalf(text_error, data.dayNumExpected, D, data.input)
			}

			if txtM != data.monthExpected {
				t.Fatalf(text_error, data.monthExpected, txtM, data.input)
			}

			if Y != data.yearExpected {
				t.Fatalf(text_error, data.yearExpected, Y, data.input)
			}

		})
	}
}
