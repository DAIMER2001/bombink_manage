package gorm

import (
	"fmt"
	"time"

	"github.com/daimer/bombink-manage/internal/domain/utils/logger"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgresDBClient(host, port, userName, password, name string) (*gorm.DB, error) {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, name)
	db, err := connectDB(gormpostgres.Open(connect))
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("NewMariaDBClient", err)
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		logger.Error("NewMariaDBClient", err)
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

func connectDB(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logger.Error("connectDB", err)
		return nil, err
	}

	return db, nil
}
