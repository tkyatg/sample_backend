package main

import (
	"log"

	"github.com/tkyatg/rental_redy_backend/adapter/env"
	"github.com/tkyatg/rental_redy_backend/adapter/http"
	"github.com/tkyatg/rental_redy_backend/adapter/sql"
)

func main() {
	env := env.NewEnv()
	dbConnection, err := sql.NewDBConnection(env.GetDBUser(), env.GetDBPassword(), env.GetDBName(), env.GetDBHost(), env.GetDBPort())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	server := http.NewServer(env.GetServerPort(), dbConnection)
	server.Serve()
}
