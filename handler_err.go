package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 300, "something went wrong")
}
