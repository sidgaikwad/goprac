package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 400, struct{}{})
}

// Path: handler_liveness.go
