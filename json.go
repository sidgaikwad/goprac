package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("the error is on the server: %v+/n", msg)
	}
	type errRespond struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errRespond{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal Json response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(dat)
	w.WriteHeader(code)

}
