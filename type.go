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

type create_user_parameters struct {
	Email string `json:"email"`
}

type create_chirp_parameters struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

type errorMessage struct {
	Error string `json:"error"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}
