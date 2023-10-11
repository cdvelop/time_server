package timeserver

import (
	"time"

	"github.com/cdvelop/timetools"
)

// AgeAt obtiene la edad de una entidad en un momento determinado.
func AgeAt(birthDate time.Time, now time.Time) int {
	// Obtener el cambio de número de año desde el nacimiento de la persona.
	age := now.Year() - birthDate.Year()

	// Si la fecha es anterior a la fecha de nacimiento, entonces no han transcurrido tantos años.

	birthDay := timetools.GetAdjustedBirthDay(birthDate.Year(), birthDate.YearDay(), now.Year(), now.YearDay())
	if now.YearDay() < birthDay {
		age -= 1
	}

	return age
}

// Edad es la abreviatura de AgeAt (birthDate, time.Now ()) y tiene el mismo uso y limitaciones.
func Age(birthDate time.Time) int {
	return AgeAt(birthDate, time.Now())
}
