package main

import (
	"encoding/json"
	"net/http"

	"github.com/misterlister/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	var params login_parameters
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	user, err := cfg.dbQueries.GetUserByEmail(req.Context(), params.Email)

	if err != nil {
		respondWithError(w, 500, DatabaseErrMsg+": "+err.Error())
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)

	if err != nil {
		respondWithError(w, 401, PasswordFailErrMsg)
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
