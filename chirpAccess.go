package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.dbQueries.GetAllChirps(req.Context())

	if err != nil {
		respondWithError(w, 500, GetChirpsErrMsg)
		return
	}

	chirpObjs := []Chirp{}

	for _, dbChirp := range chirps {
		newChirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		}
		chirpObjs = append(chirpObjs, newChirp)
	}

	respondWithJSON(w, 200, chirpObjs)
}

func (cfg *apiConfig) handlerGetChirpByID(w http.ResponseWriter, req *http.Request) {
	chirpID := req.PathValue(ChirpID)

	chirpUUID, err := uuid.Parse(chirpID)

	if err != nil {
		respondWithError(w, 400, UUIDErrMsg)
		return
	}

	dbChirp, err := cfg.dbQueries.GetChirpByID(req.Context(), chirpUUID)

	if err != nil {
		respondWithError(w, 404, NoChirpFoundErrMsg)
		return
	}

	chirpObj := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	}

	respondWithJSON(w, 200, chirpObj)
}
