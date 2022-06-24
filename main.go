package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app"
)

func main() {
	fmt.Println("Starting HTTP server for scorecard-bigquery-auth on port 6767...")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", app.Index)
	r.HandleFunc("/query/arbitary", app.PostResultsHandler).Methods(http.MethodPost)
	r.HandleFunc("/query/vulnerabilities", app.PostResultsHandler).Methods(http.MethodPost)
	r.HandleFunc("/query/dependencies", app.PostResultsHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
