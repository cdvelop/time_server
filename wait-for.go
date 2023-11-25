package timeserver

import "time"

// WaitFor espera el número especificado de milisegundos y luego ejecuta la función de retorno de llamada.
func (timeServer) WaitFor(milliseconds int, callback func()) {
	duration := time.Duration(milliseconds) * time.Millisecond
	time.Sleep(duration)
	callback()
}
