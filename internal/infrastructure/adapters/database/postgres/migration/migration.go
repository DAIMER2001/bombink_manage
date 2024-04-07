package migration

import (
	"os"

	"github.com/daimer/bombink-manage/internal/domain/utils/logger"
	"github.com/daimer/bombink-manage/internal/infrastructure/adapters/database/postgres/models"
	gormPkg "github.com/daimer/bombink-manage/internal/infrastructure/pkg/gorm"
)

func StartMigration() error {
	conn, err := gormPkg.NewPostgresDBClient(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		logger.Error("LoadMigration", err)
		return err
	}

	err = conn.Debug().AutoMigrate(
		&models.Company{},
		&models.DataList{},
		&models.Dropshipping{},
		&models.DropshippingCompany{},
		&models.GalleryTattoo{},
		&models.Product{},
		&models.ProductsFile{},
		&models.ProductsStock{},
		&models.Sale{},
		&models.Seller{},
		&models.Service{},
		&models.TattooArtist{},
		&models.TattooReactions{},
		&models.Tattoos{},
		&models.TypeList{},
		&models.User{},
	)
	if err != nil {
		logger.Error("LoadMigration", err)
		return err
	}

	return nil
}
