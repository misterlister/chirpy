package main

import (
	"encoding/json"
	"net/http"

	"github.com/misterlister/chirpy/internal/database"
)

func (cfg *apiConfig) handlerPostChirp(w http.ResponseWriter, req *http.Request) {
	var params create_chirp_parameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	if len(params.Body) > MaxChirpLength {
		respondWithError(w, 400, TooLongErrMsg)
		return
	}

	cleanedBody := removeBadWords(params.Body)

	cleanedParams := database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: params.UserID,
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
		UserID:    cleanedChirp.UserID,
	}

	respondWithJSON(w, 201, chirpObj)
}
