package sql

import (
	"database/sql"

	// qg
	_ "github.com/lib/pq"
)

// NewSQLConnect func
func NewDBConnection(dbUser string, dbPassword string, dbName string, dbHost string, dbport string) (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://"+dbUser+":"+dbPassword+"@"+dbHost+":"+dbport+"/"+dbName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
