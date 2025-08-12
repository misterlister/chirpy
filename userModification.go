package main

import (
	"encoding/json"
	"net/http"

	"github.com/misterlister/chirpy/internal/auth"
	"github.com/misterlister/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUserCreate(w http.ResponseWriter, req *http.Request) {
	var params createUserParameters
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

	hashedPassword, err := auth.HashPassword(params.Password)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	userParams := database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
	}

	user, err := cfg.dbQueries.CreateUser(req.Context(), userParams)

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

func (cfg *apiConfig) handlerUserUpdate(w http.ResponseWriter, req *http.Request) {
	uuid, err := cfg.validateLoginStatus(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	var params userUpdateParameters
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	userParams := database.UpdateUserEmailAndPasswordParams{
		ID:             uuid,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	}

	user, err := cfg.dbQueries.UpdateUserEmailAndPassword(req.Context(), userParams)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	userObj := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respondWithJSON(w, 200, userObj)
}
