package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting HTTP server for scorecard-bigquery-auth on port %s...\n", port)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", app.Index)
	r.HandleFunc("/query/{type}", app.GetResultsHandler).Methods(http.MethodGet)
	r.HandleFunc("/query/{type}", app.PostResultsHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
