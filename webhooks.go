package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/misterlister/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, req *http.Request) {
	var params polkaWebhookParameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	validated, err := cfg.validateApiKey(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	if !validated {
		w.WriteHeader(401)
		return
	}

	userUuid, err := uuid.Parse(params.Data.UserId)

	if err != nil {
		respondWithError(w, 404, err.Error())
		return
	}

	if params.Event == UpgradeUser {
		err = cfg.dbQueries.UpgradeToChirpyRed(req.Context(), userUuid)

		if err != nil {
			respondWithError(w, 404, err.Error())
			return
		}
	}

	w.WriteHeader(204)
}

func (cfg *apiConfig) validateApiKey(header http.Header) (bool, error) {
	apiKey, err := auth.GetAPIKey(header)

	if err != nil {
		return false, err
	}

	if apiKey != cfg.polkaKey {
		return false, nil
	}

	return true, nil
}
