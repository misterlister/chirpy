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
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("Error creating database connection")
		return
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		dbQueries: dbQueries,
	}

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
