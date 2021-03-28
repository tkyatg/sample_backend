package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Message string
}

func main() {
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
