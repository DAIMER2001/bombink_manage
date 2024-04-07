package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModelError struct {
	Msn        string
	StatusCode int
}

var typeErrors = map[string]ModelError{
	"ErrorLenName":          {Msn: "Error tamaño de nombre no es valido", StatusCode: http.StatusNotFound},
	"ErrorLenVersion":       {Msn: "Error tamaño de versión no es valido", StatusCode: http.StatusNotFound},
	"ErrorLenCountry":       {Msn: "Error tamaño de nombre de país no es valido", StatusCode: http.StatusNotFound},
	"ErrorIdCompanyInvalid": {Msn: "Error id company numero invalido", StatusCode: http.StatusNotFound},
	"ErrorStatusInvalid":    {Msn: "Error stattus no es valido", StatusCode: http.StatusNotFound},
	"ErrorTimerInvalid":     {Msn: "Error timer no es valido", StatusCode: http.StatusNotFound},
	"ErrorOrderInvalid":     {Msn: "Error order no es valido", StatusCode: http.StatusNotFound},
	"ErrorLenPage":          {Msn: "Error tamaño pagina fila actual", StatusCode: http.StatusNotFound},
	"ErrorLenPageSize":      {Msn: "Error tamaño de paginación", StatusCode: http.StatusNotFound},
	"ErrorLenSearch":        {Msn: "Error tamaño busqueda fila actual", StatusCode: http.StatusNotFound},
	"Default":               {Msn: "Ocurrio un error inexperado", StatusCode: http.StatusInternalServerError},
	"IDUserEmpty":           {Msn: "El id del usuario vacío", StatusCode: http.StatusBadRequest},
}

func OrchestratorResponsesWithError(c echo.Context, err error) error {
	if err != nil {
		value, ok := typeErrors[err.Error()]
		if ok {
			return BuildResponseDto(c, value.StatusCode, value.Msn, err.Error())
		}
	}
	return BuildResponseDto(c, typeErrors["Default"].StatusCode, typeErrors["Default"].Msn, "")
}
