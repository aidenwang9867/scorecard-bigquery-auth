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
		system := mux.Vars(r)["system"]
		name := mux.Vars(r)["name"]
		version := mux.Vars(r)["version"]
		// The Dependency.Name field is required, none indicating the input is not valid.
		if name == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "error unmarshaling input JSON")
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
