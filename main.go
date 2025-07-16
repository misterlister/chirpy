package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const appPref = "/app"
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
	serveMux.Handle(appPref+"/", http.StripPrefix(appPref, http.FileServer(http.Dir("."))))
	serveMux.Handle("/healthz", http.HandlerFunc(handlerReadiness))
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
