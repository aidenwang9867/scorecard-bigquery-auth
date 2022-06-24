package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	endpts := struct {
		QueryVulnerabilities string `json:"query_vuln"`
		QueryDependencies    string `json:"query_deps"`
		QueryArbitary        string `json:query_arbitary`
	}{
		QueryVulnerabilities: "/query/vulnerabilities",
		QueryDependencies:    "/query/dependencies",
		QueryArbitary:        "/query/arbitary",
	}
	endptsBytes, err := json.MarshalIndent(endpts, "", " ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprint(w, string(endptsBytes)); err != nil {
		log.Fatal(err)
	}
}
