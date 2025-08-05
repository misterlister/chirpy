package main

import (
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	params, err := decodeParams(req)

	if err != nil {
		respondWithError(w, 400, UnknownErrMsg)
		return
	}

	if len(params.Body) > MaxChirpLength {
		respondWithError(w, 400, TooLongErrMsg)
		return
	}

	cleaned_body := removeBadWords(params.Body)

	resp := validMessage{
		Body: cleaned_body,
	}

	respondWithJSON(w, http.StatusOK, resp)
}
