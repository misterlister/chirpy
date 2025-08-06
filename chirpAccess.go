package main

import "net/http"

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.dbQueries.GetAllChirps(req.Context())

	if err != nil {
		respondWithError(w, 500, GetChirpsErrMsg)
		return
	}

	chirpObjs := []Chirp{}

	for _, dbChirp := range chirps {
		newChirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		}
		chirpObjs = append(chirpObjs, newChirp)
	}

	respondWithJSON(w, 200, chirpObjs)
}
