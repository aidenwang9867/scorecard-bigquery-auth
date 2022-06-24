package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	endpts := struct {
		GetRepoResults string `json:"get_repo_results"`
		GetRepoBadge   string `json:"get_repo_badge"`
	}{
		GetRepoResults: "/projects/{host}/{owner}/{repository}",
		GetRepoBadge:   "/projects/{host}/{owner}/{repository}/badge",
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
