package env

import (
	"os"

	"github.com/tkyatg/rental_redy_backend/shared"
)

type (
	environment struct {
		dbHost     string
		dbPort     string
		dbUser     string
		dbPassword string
		dbName     string
		serverPort string
	}
)

const defaultPort = "8080"

// NewEnv func
func NewEnv() shared.Env {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}

	return &environment{
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		serverPort,
	}
}

func (t *environment) GetDBHost() string {
	return t.dbHost
}
func (t *environment) GetDBPort() string {
	return t.dbPort
}
func (t *environment) GetDBUser() string {
	return t.dbUser
}
func (t *environment) GetDBPassword() string {
	return t.dbPassword
}
func (t *environment) GetDBName() string {
	return t.dbName
}
func (t *environment) GetServerPort() string {
	return t.serverPort
}
