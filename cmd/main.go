package main

import (
	"os"

	"github.com/daimer/bombink-manage/internal/domain/utils/logger"
	"github.com/daimer/bombink-manage/internal/infrastructure/adapters/database/postgres/migration"
	"github.com/daimer/bombink-manage/internal/infrastructure/enums"
	apirestfactory "github.com/daimer/bombink-manage/internal/infrastructure/factory/apiRestFactory"
	apiRestEnum "github.com/daimer/bombink-manage/internal/infrastructure/factory/apiRestFactory/enums"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", string(enums.EnvLocal))
	}

	// Carga el .env solo si estamos en entorno local
	if os.Getenv("ENV") == string(enums.EnvLocal) {
		err := godotenv.Load()
		if err != nil {
			logger.Error("Main", err)
		}
	}

	bootstrap()
}

func bootstrap() {

	err := migration.StartMigration()
	if err != nil {
		panic(err)
	}

	apiRest := apirestfactory.NewAPIRest(apiRestEnum.Echo)
	apiRest.RunServer()
}
