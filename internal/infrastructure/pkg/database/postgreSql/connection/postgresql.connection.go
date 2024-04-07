package postgreSql

import (
	"errors"
	"fmt"
	"os"

	"github.com/daimer/bombink-manage/internal/domain/constants"
	secret "github.com/daimer/bombink-manage/internal/infrastructure/pkg/secret/postgresql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	connectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func NewPostgresConection(credentials secret.DBCredentials) (db *gorm.DB) {
	if credentials.Username != "" || credentials.Password != "" {
		dsnco := fmt.Sprintf(connectionString, credentials.Host, credentials.Port, credentials.Username, credentials.Password, credentials.Name)
		if db, err := gorm.Open(gormpostgres.Open(dsnco), &gorm.Config{
			SkipDefaultTransaction: true,
		}); err != nil {
			log.Println("Error en la conexion a la bd: ", err.Error())
			panic(err)
		} else {
			return defineMultipleDatabaseConnections(db, dsnco)
		}
	}
	return
}

func defineMultipleDatabaseConnections(db *gorm.DB, dnsco string) *gorm.DB {
	err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{gormpostgres.Open(dnsco)},
	}, constants.CODBConnectionName))
	if err != nil {
		log.Errorf("Error en la multi conexion a la bd: %v", err.Error())
		panic(err)
	}
	log.Printf("Conexion a la bd exitosa en ambiente %v", os.Getenv("ENV"))
	return db
}

func GetDatabaseNameConnection(countryCode string) (string, error) {
	var databaseConnectionName string
	switch countryCode {
	case constants.CountryCodeCO:
		databaseConnectionName = constants.CODBConnectionName
	default:
		return "", errors.New("error database not available")
	}
	return databaseConnectionName, nil
}
