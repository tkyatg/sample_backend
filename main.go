package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tkyatg/rental_redy_backend/adapter/env"
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

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ping := Result{Message: "ok"}

	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
