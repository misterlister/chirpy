package main

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
}

type parameters struct {
	Body  string `json:"body"`
	Email string `json:"email"`
}

type errorMessage struct {
	Error string `json:"error"`
}

type validMessage struct {
	Body string `json:"cleaned_body"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}
