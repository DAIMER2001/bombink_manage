package secret

import (
	"os"
)

type DBCredentials struct {
	Username string `json:"DB_USER"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	Host     string `json:"DB_HOST"`
	Name     string `json:"DB_NAME"`
}

type TokenCredentials struct {
	BombinkSecret string `json:"BOMBINK_TOKEN_SECRET"`
}

func GetDBCredentials() DBCredentials {
	return DBCredentials{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
	}

}
