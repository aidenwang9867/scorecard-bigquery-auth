package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app/query"
	"github.com/gorilla/mux"
)

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	queryType := mux.Vars(r)["type"]
	switch queryType {
	case "vulnerabilities", "dependencies":
		system := r.URL.Query().Get("system")
		name := r.URL.Query().Get("name")
		version := r.URL.Query().Get("version")
		// The Dependency.Name field is required, none indicating the input is not valid.
		if system == "" || name == "" || version == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "error found in query parameters")
			if err != nil {
				log.Printf("error during Write: %v", err)
			}
			return
		}
		d := query.Dependency{
			Ecosystem: system,
			Name:      name,
			Version:   version,
		}
		if queryType == "vulnerabilities" {
			helperQueryVulnerabilities(w, d)
		} else if queryType == "dependencies" {
			helperQueryDependencies(w, d)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf("query type %s not supported", queryType)
	}
}
