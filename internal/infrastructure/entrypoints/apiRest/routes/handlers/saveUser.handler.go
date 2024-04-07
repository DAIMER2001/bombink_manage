package handlers

import (
	"errors"
	"net/http"

	useCase "github.com/daimer/bombink-manage/internal/application/usecases/saveUser"
	usecasedto "github.com/daimer/bombink-manage/internal/application/usecases/saveUser/dtos"
	"github.com/daimer/bombink-manage/internal/infrastructure/deps"
	api "github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/response"
	"github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes/dtos"
	"github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes/utils"
	"github.com/labstack/echo/v4"
)

// @Summary Service to activate missions
// @Description Service to activate missions
// @Tags ActivateUser
// @Accept json
// @Produce json
// @Param data body dtos.BombinkResponse true "Save user"
// @Success 201 {object} dtos.BombinkResponse "user guardado con éxito"
// @Failure 500 {object} api.RetomaRegistroResponseErrorInexperado "Cuando ocurre un error inexperado"
// @Router /commerce/manage-bombink/user/create [post]
func SaveUserHandler(c echo.Context) error {
	var body = &dtos.UserDto{}
	if err := c.Bind(body); err != nil {
		return utils.OrchestratorResponsesWithError(c, errors.New("InvalidBodyFormat"))
	}

	misteryShopperDB, err := deps.NewBombinkDBAdapter()
	if err != nil {
		return utils.OrchestratorResponsesWithError(c, err)
	}

	data := convertBodySaveUser(body)
	because := useCase.New(misteryShopperDB)
	idUser, err := because.SaveUser(data)
	if err != nil {
		return utils.OrchestratorResponsesWithError(c, err)
	}

	return c.JSON(http.StatusOK, dtos.BombinkResponse{
		Message: "user creado con éxito",
		Data: api.SaveUserResponse{
			UID:     idUser,
			Success: true,
		},
	})
}

func convertBodySaveUser(body *dtos.UserDto) *usecasedto.UserDto {
	return &usecasedto.UserDto{
		UID:         body.UID,
		Phone:       body.Phone,
		Password:    body.Password,
		Birthday:    body.Birthday,
		IsSuperuser: body.IsSuperuser,
		Username:    body.Username,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		IsStaff:     body.IsStaff,
		IsActive:    body.IsActive,
	}
}
