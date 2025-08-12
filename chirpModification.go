package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/misterlister/chirpy/internal/database"
)

func (cfg *apiConfig) handlerPostChirp(w http.ResponseWriter, req *http.Request) {
	var params createChirpParameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	userUuid, err := cfg.validateLoginStatus(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	if len(params.Body) > MaxChirpLength {
		respondWithError(w, 400, TooLongErrMsg)
		return
	}

	cleanedBody := removeBadWords(params.Body)

	cleanedParams := database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: userUuid,
	}

	cleanedChirp, err := cfg.dbQueries.CreateChirp(req.Context(), cleanedParams)

	if err != nil {
		respondWithError(w, 500, DatabaseErrMsg+": "+err.Error())
		return
	}

	chirpObj := Chirp{
		ID:        cleanedChirp.ID,
		CreatedAt: cleanedChirp.CreatedAt,
		UpdatedAt: cleanedChirp.UpdatedAt,
		Body:      cleanedChirp.Body,
		UserID:    userUuid,
	}

	respondWithJSON(w, 201, chirpObj)
}

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, req *http.Request) {
	userUuid, err := cfg.validateLoginStatus(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	chirpID := req.PathValue(ChirpID)

	chirpUUID, err := uuid.Parse(chirpID)

	if err != nil {
		respondWithError(w, 400, UUIDErrMsg)
		return
	}

	dbChirp, err := cfg.dbQueries.GetChirpByID(req.Context(), chirpUUID)

	if err != nil {
		respondWithError(w, 404, err.Error())
		return
	}

	if dbChirp.UserID != userUuid {
		respondWithError(w, 403, DeleteAuthErrMsg)
		return
	}

	err = cfg.dbQueries.DeleteChirpByID(req.Context(), chirpUUID)

	if err != nil {
		respondWithError(w, 404, err.Error())
		return
	}
	w.WriteHeader(204)
}
