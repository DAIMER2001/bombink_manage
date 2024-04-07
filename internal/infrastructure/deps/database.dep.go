package deps

import (
	"os"

	"github.com/daimer/bombink-manage/internal/infrastructure/adapters/database"
	gormPkg "github.com/daimer/bombink-manage/internal/infrastructure/pkg/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbError bool

func NewBombinkDBAdapter() (*database.BombinkDB, error) {
	var err error
	if db == nil {
		err = connectDB()
		if err != nil {
			dbError = true
			return nil, err
		}
	}

	if dbError {
		err = connectDB()
		if err != nil {
			dbError = true
			return nil, err
		}
	}

	dbError = false
	adapter := database.NewBombinkDB(db)
	return adapter, nil
}

func connectDB() error {
	var err error
	db, err = gormPkg.NewPostgresDBClient(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	return err
}
