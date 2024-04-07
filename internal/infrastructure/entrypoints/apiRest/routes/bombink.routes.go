package routes

import (
	_ "github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes/docs"
	"github.com/daimer/bombink-manage/internal/infrastructure/entrypoints/apiRest/routes/handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Bombink End Points
// @version 1.0
// @description Service to connect users
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /commerce/manage-bombink
// @schemes http
func AdminBombinkRoutes(group *echo.Group) {
	group.GET("/healthz", handlers.HealthCheckHandler)
	group.POST("/user/create", handlers.SaveUserHandler)
	group.GET("/swagger/*", echoSwagger.WrapHandler)
}
