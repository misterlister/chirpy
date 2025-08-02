package main

import (
	"sync/atomic"

	"github.com/misterlister/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
}

type parameters struct {
	Body string `json:"body"`
}

type errorMessage struct {
	Error string `json:"error"`
}

type validMessage struct {
	Body string `json:"cleaned_body"`
}
