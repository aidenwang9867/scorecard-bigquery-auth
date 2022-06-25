package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
		if err != nil || d.Name == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprint(w, "error unmarshaling input JSON")
			if err != nil {
				log.Printf("error during Write: %v", err)
			}
			return
		}
		if queryType == "vulnerabilities" {
			helperQueryVulnerabilities(w, d)
			return
		} else if queryType == "dependencies" {
			helperQueryDependencies(w, d)
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		log.Printf("query type %s not supported", queryType)
		return
	}
}

func helperQueryVulnerabilities(w http.ResponseWriter, d query.Dependency) {
	auth, err := query.Authenticate()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error authenticating the service account")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
	}
	vuln, err := query.GetVulnerabilitiesBySystemNameVersion(
		auth,
		strings.ToUpper(d.Ecosystem),
		d.Name,
		d.Version,
	)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error retrieving vuln by sys/name/version")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vuln)
}

func helperQueryDependencies(w http.ResponseWriter, d query.Dependency) {
	auth, err := query.Authenticate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error authenticating the service account")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
	}
	deps, err := query.GetDependenciesBySystemNameVersion(
		auth,
		strings.ToUpper(d.Ecosystem),
		d.Name,
		d.Version,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error retrieving vuln by sys/name/version")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(deps)
}
