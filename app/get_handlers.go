package app

import (
	"errors"
	"log"
	"net/http"
)

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("get request not supported")
	http.Error(w, err.Error(), http.StatusNotImplemented)
	log.Printf("err: %v", err)
}
