package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	errMessage := errorMessage{
		Error: msg,
	}

	errMessageReturn, err := json.Marshal(errMessage)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("error: unable to marshal valid response message")
		return
	}
	w.Header().Set(ContentType, JsonHeader)
	w.WriteHeader(code)
	w.Write(errMessageReturn)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error marshalling JSON: %s", err)
		return
	}
	w.Header().Set(ContentType, JsonHeader)
	w.WriteHeader(code)
	w.Write(dat)
}
