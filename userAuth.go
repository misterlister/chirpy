package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	var params loginParameters
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

	expirationDuration := DefaultLoginExpiry

	if params.ExpiresInSeconds != nil && *params.ExpiresInSeconds <= DefaultLoginExpiry && *params.ExpiresInSeconds > 0 {
		expirationDuration = *params.ExpiresInSeconds
	}

	newToken, err := auth.MakeJWT(user.ID, cfg.secret, time.Duration(expirationDuration)*time.Second)

	if err != nil {
		respondWithError(w, 500, TokenFailErrMsg+": "+err.Error())
		return
	}

	loginResp := loginResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     newToken,
	}

	respondWithJSON(w, 200, loginResp)
}

func (cfg *apiConfig) validateLoginStatus(header http.Header) (uuid.UUID, error) {
	token, err := auth.GetBearerToken(header)

	if err != nil {
		return uuid.Nil, err
	}

	uuidResponse, err := auth.ValidateJWT(token, cfg.secret)

	if err != nil {
		return uuid.Nil, err
	}

	return uuidResponse, nil
}
