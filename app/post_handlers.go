package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app/query"
	"github.com/gorilla/mux"
)

func PostResultsHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "error reading request body")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	queryType := mux.Vars(r)["type"]
	switch queryType {
	case "vulnerabilities", "dependencies":
		d := query.Dependency{}
		err := json.Unmarshal(reqBody, &d)
		// The Dependency.Name field is required, none indicating the input is not valid.
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "error unmarshaling input JSON")
			if err != nil {
				log.Printf("error during Write: %v", err)
			}
			return
		}
		if d.Ecosystem == "" || d.Name == "" || d.Version == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "error found in post body")
			if err != nil {
				log.Printf("error during Write: %v", err)
			}
			return
		}
		if queryType == "vulnerabilities" {
			helperQueryVulnerabilities(w, d)
		} else if queryType == "dependencies" {
			helperQueryDependencies(w, d)
		}
	case "arbitrary":
		query := string(reqBody)
		if query == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "empty query received")
			if err != nil {
				log.Printf("error during Write: %v", err)
			}
			return
		}
		helperArbitararyQuery(w, query)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf("query type %s not supported", queryType)
	}
}
