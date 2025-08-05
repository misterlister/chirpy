package main

import (
	"encoding/json"
	"net/http"
)

func decodeParams(req *http.Request) (parameters, error) {
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	return params, err
}
