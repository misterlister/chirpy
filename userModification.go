package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUserCreate(w http.ResponseWriter, req *http.Request) {
	var params create_user_parameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	if !(len(params.Email) > 0) {
		respondWithError(w, 400, MissingParamErrMsg)
		return
	}

	user, err := cfg.dbQueries.CreateUser(req.Context(), params.Email)

	if err != nil {
		respondWithError(w, 500, DatabaseErrMsg+": "+err.Error())
		return
	}

	userObj := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respondWithJSON(w, 201, userObj)
}
