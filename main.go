package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting HTTP server for scorecard-bigquery-auth on port 6767...")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", app.Index)
	r.HandleFunc("/projects/{host}/{orgName}/{repoName}", app.PostResultsHandler).Methods(http.MethodPost)
	r.HandleFunc("/projects/{host}/{orgName}/{repoName}", app.GetResultsHandler).Methods(http.MethodGet)
	r.HandleFunc("/projects/{host}/{orgName}/{repoName}/badge", app.GetBadgeHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
