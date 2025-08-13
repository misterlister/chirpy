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
	secret := os.Getenv("SECRET")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(DatabaseInitErrMsg)
		return
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		dbQueries: dbQueries,
		platform:  platform,
		secret:    secret,
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
	serveMux.HandleFunc(GetReq+ApiPrefix+ChirpsPath, apiCfg.handlerGetAllChirps)
	serveMux.HandleFunc(GetReq+ApiPrefix+ChirpsPath+"/{"+ChirpID+"}", apiCfg.handlerGetChirpByID)

	// POST requests
	serveMux.HandleFunc(PostReq+AdminPrefix+ResetPath, apiCfg.handlerReset)
	serveMux.HandleFunc(PostReq+ApiPrefix+ChirpsPath, apiCfg.handlerPostChirp)
	serveMux.HandleFunc(PostReq+ApiPrefix+UsersPath, apiCfg.handlerUserCreate)
	serveMux.HandleFunc(PostReq+ApiPrefix+LoginPath, apiCfg.handlerLogin)
	serveMux.HandleFunc(PostReq+ApiPrefix+RefreshPath, apiCfg.handlerRefresh)
	serveMux.HandleFunc(PostReq+ApiPrefix+RevokePath, apiCfg.handlerRevoke)
	serveMux.HandleFunc(PostReq+ApiPrefix+PolkaPath+WebhooksPath, apiCfg.handlerPolkaWebhook)

	// PUT requests
	serveMux.HandleFunc(PutReq+ApiPrefix+UsersPath, apiCfg.handlerUserUpdate)

	// DELETE requests
	serveMux.HandleFunc(DeleteReq+ApiPrefix+ChirpsPath+"/{"+ChirpID+"}", apiCfg.handlerDeleteChirp)

	log.Printf("Serving on port: %s\n", Port)
	log.Fatal(server.ListenAndServe())
}
