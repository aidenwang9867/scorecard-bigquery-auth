package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostResultsHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "error reading request body")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	fmt.Println(reqBody)
}
