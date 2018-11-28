package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func characterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, id)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home")
}
