package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Printf("Responding with 5XX Server error: %s", message)
	}

	respondWithJSON(w, statusCode, ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonPayload)
}
