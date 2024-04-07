package utils

import (
	"github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes/dtos"
	"github.com/labstack/echo/v4"
)

func BuildResponseDto(c echo.Context, httpStatus int, msg, msnError string) error {
	return c.JSON(httpStatus, dtos.BombinkResponse{
		Message: msg,
		Data: map[string]string{
			"code": msnError,
		},
	})
}
