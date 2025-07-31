package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	if len(params.Body) > MaxChirpLength {
		respondWithError(w, 400, TooLongErrMsg)
		return
	}

	resp := validMessage{
		Valid: true,
	}

	respondWithJSON(w, http.StatusOK, resp)
}
