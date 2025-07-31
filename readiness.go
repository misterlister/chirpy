package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(ContentType, TextHeader)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
