package main

import "github.com/google/uuid"

type createUserParameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createChirpParameters struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

type loginParameters struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
