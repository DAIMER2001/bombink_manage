package migration

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	secret "github.com/daimer/bombink-manage/internal/infrastructure/pkg/secret/postgresql"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	connectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func StartDBMigration(credentials secret.DBCredentials) {
	if credentials.Username == "" || credentials.Password == "" {
		pqDNS := fmt.Sprintf(connectionString, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME_CO"))
		buildMigration(pqDNS)
	}
}

func buildMigration(pqDNS string) {
	db, err := sql.Open("postgres", pqDNS)
	if err != nil {
		panic("error opening connection to run migration: " + err.Error())
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic("error initializing the client in charge of executing the migration: " + err.Error())
	}
	defer driver.Close()
	m, err := migrate.NewWithDatabaseInstance("file://db/migration/DDL", "postgres", driver)
	if err != nil {
		panic("error connecting client with database to execute migration: " + err.Error())
	}
	err = m.Up()
	if err != nil {
		log.Info("Migration status: " + err.Error())
	}
}
