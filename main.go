package main

import (
	"log"

	"github.com/tkyatg/rental_redy_backend/adapter/env"
	"github.com/tkyatg/rental_redy_backend/adapter/http"
	"github.com/tkyatg/rental_redy_backend/adapter/sql"
)

type Result struct {
	Message string
}

func main() {
	env := env.NewEnv()
	dbConnection, err := sql.NewGormConnection(env.GetDBUser(), env.GetDBPassword(), env.GetDBName(), env.GetDBHost(), env.GetDBPort())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	server := http.NewServer(env.GetServerPort(), dbConnection)
	server.Serve()
}
