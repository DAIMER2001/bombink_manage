package logger

import (
	"log"
	"runtime"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

/*
	Info:

Función que crea un log con el formato:
  - nombre_archivo:numero_linea type="info" name="functionName" data="message"

parameters:
  - functionName : Nombre de la función desde la que se llama el log
  - message : Mensaje que se quiere imprimir
*/
func Info(functionName string, message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf(Green+"%s:%d type=\"info\" name=\"%s\" data=\"%v\""+Reset, file, line, functionName, message)
}

/*
	Error:

Función que crea un log con el formato:
  - nombre_archivo:numero_linea type="error" name="functionName" data="message"

parameters:
  - functionName : Nombre de la función desde la que se llama el log
  - err : instancia del error que se quiere imprimir
*/
func Error(functionName string, err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf(Red+"%s:%d type=\"error\" name=\"%s\" data=\"%v\""+Reset, file, line, functionName, err)
}
