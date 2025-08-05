package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, UnauthorizedErrMsg)
		return
	}

	err := cfg.dbQueries.DeleteAllUsers(req.Context())

	if err != nil {
		respondWithError(w, 500, DatabaseErrMsg)
		return
	}

	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ResetMsg))
}
