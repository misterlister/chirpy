package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/misterlister/chirpy/internal/database"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(DatabaseInitErrMsg)
		return
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		dbQueries: dbQueries,
		platform:  platform,
	}

	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + Port,
		Handler: serveMux,
	}

	// General handler
	serveMux.Handle(AppPrefix+"/", apiCfg.middlewareMetricsInc(http.StripPrefix(AppPrefix, http.FileServer(http.Dir(".")))))

	// GET requests
	serveMux.HandleFunc(GetReq+ApiPrefix+HealthPath, handlerReadiness)
	serveMux.HandleFunc(GetReq+AdminPrefix+MetricPath, apiCfg.handlerMetrics)

	// POST requests
	serveMux.HandleFunc(PostReq+AdminPrefix+ResetPath, apiCfg.handlerReset)
	serveMux.HandleFunc(PostReq+ApiPrefix+ChirpsPath, apiCfg.handlerPostChirp)
	serveMux.HandleFunc(PostReq+ApiPrefix+UsersPath, apiCfg.handlerUserCreate)

	log.Printf("Serving on port: %s\n", Port)
	log.Fatal(server.ListenAndServe())
}
