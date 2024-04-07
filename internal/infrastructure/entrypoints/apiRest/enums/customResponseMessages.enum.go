package enums

type CustomResponseMessages string

const (
	DatosEnviadosInvalidos CustomResponseMessages = "Datos enviados inválidos"
	ErrorInesperado        CustomResponseMessages = "Ocurrio un error inesperado"
	OK                     CustomResponseMessages = "Ok"
)
