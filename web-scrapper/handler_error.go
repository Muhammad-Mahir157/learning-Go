package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "An error occurred while processing your request. Please try again later.")
}
