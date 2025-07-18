package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

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
	log.Printf("Serving on port: %s\n", Port)
	log.Fatal(server.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(ContentType, TextHeader)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	hits := cfg.fileserverHits.Load()
	w.Header().Set(ContentType, HtmlHeader)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		fmt.Sprintf(
			`<html>
				<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
				</body>
			</html>`, hits)))
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
	return handler
}
