package main

import (
	"log"
	"net/http"
)

func main() {
	var apiCfg apiConfig
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + Port,
		Handler: serveMux,
	}
	serveMux.Handle(AppPrefix+"/", apiCfg.middlewareMetricsInc(http.StripPrefix(AppPrefix, http.FileServer(http.Dir(".")))))
	serveMux.HandleFunc(GetReq+ApiPrefix+HealthPath, handlerReadiness)
	serveMux.HandleFunc(GetReq+AdminPrefix+MetricPath, apiCfg.handlerMetrics)
	serveMux.HandleFunc(PostReq+AdminPrefix+ResetPath, apiCfg.handlerReset)
	serveMux.HandleFunc(PostReq+ApiPrefix+ValidateChirpPath, handlerValidateChirp)
	log.Printf("Serving on port: %s\n", Port)
	log.Fatal(server.ListenAndServe())
}
