package apirestfactory

import (
	"os"

	apirest "github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest"
	"github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes"
	"github.com/daimer/bombink-manage/internal/infrastructure/factory/apiRestFactory/enums"
	"github.com/daimer/bombink-manage/internal/infrastructure/interfaces"
)

func NewAPIRest(provider enums.APIRestProviders) interfaces.IAPIRest {
	switch provider {
	case enums.Echo:
		return apirest.NewAPIRestServer(
			os.Getenv("API_REST_PORT"),
			os.Getenv("API_REST_BASE_URL"),
			routes.AdminBombinkRoutes,
		)
	default:
		return nil
	}
}
