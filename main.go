package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app"
)

const (
	PORT = "8080"
)

func main() {

	fmt.Printf("Starting HTTP server for scorecard-bigquery-auth on port %s...\n", PORT)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", app.Index)
	r.HandleFunc("/query/{type}", app.GetResultsHandler).Methods(http.MethodGet)
	r.HandleFunc("/query/{type}", app.PostResultsHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
