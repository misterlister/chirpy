package main

import (
	"net/http"
	"time"

	"github.com/misterlister/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	user, err := cfg.dbQueries.GetUserFromRefreshToken(req.Context(), refreshToken)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.secret, time.Hour)

	if err != nil {
		respondWithError(w, 500, TokenFailErrMsg+": "+err.Error())
		return
	}

	response := refreshResponse{
		Token: accessToken,
	}

	respondWithJSON(w, 200, response)
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	err = cfg.dbQueries.RevokeRefreshToken(req.Context(), refreshToken)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	w.WriteHeader(204)
}
