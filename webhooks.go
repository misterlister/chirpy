package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, req *http.Request) {
	var params polkaWebhookParameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, err.Error())
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
