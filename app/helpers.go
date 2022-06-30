package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aidenwang9867/scorecard-bigquery-auth/app/query"
)

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

func helperArbitararyQuery(w http.ResponseWriter, q string) {
	auth, err := query.Authenticate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error authenticating the service account")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
	}
	results, err := query.GetResultsByArbitraryQuery(
		auth,
		q,
	)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "error retrieving query results")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(results))
}
