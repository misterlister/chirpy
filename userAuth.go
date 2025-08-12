package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/chirpy/internal/auth"
	"github.com/misterlister/chirpy/internal/database"
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

	accessToken, err := auth.MakeJWT(user.ID, cfg.secret, time.Hour)

	if err != nil {
		respondWithError(w, 500, TokenFailErrMsg+": "+err.Error())
		return
	}

	var refreshTokenParams database.CreateRefreshTokenParams

	refreshToken, _ := auth.MakeRefreshToken()

	refreshTokenParams.Token = refreshToken
	refreshTokenParams.UserID = user.ID

	_, err = cfg.dbQueries.CreateRefreshToken(req.Context(), refreshTokenParams)

	loginResp := loginResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
