package main

import "sync/atomic"

type apiConfig struct {
	fileserverHits atomic.Int32
}

type parameters struct {
	Body string `json:"body"`
}

type errorMessage struct {
	Error string `json:"error"`
}

type validMessage struct {
	Valid bool `json:"valid"`
}
