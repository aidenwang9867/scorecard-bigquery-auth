package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(
		"GET request to route %v not supported, please use POST with a valid request body",
		r.RequestURI,
	)
	err := errors.New(msg)
	http.Error(w, err.Error(), http.StatusNotImplemented)
	log.Printf("err: %v", err)
}
