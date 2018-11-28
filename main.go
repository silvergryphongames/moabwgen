package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/character/{id:[0-9]+}", characterHandler)

	http.Handle("/", r)
}
