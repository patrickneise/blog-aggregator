package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/patrickneise/blog-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiConfig := apiConfig{
		DB: dbQueries,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/err", apiConfig.handlerErr)
	mux.HandleFunc("GET /v1/healthz", apiConfig.handlerHealhtz)

	mux.HandleFunc("GET /v1/feeds", apiConfig.handlerFeedsGetAll)
	mux.Handle("POST /v1/feeds", apiConfig.middlewareAuth(apiConfig.handlerFeedsCreate))

	mux.Handle("GET /v1/feed_follows/", apiConfig.middlewareAuth(apiConfig.handlerFeedFollowsGet))
	mux.Handle("POST /v1/feed_follows/", apiConfig.middlewareAuth(apiConfig.handlerFeedFollowsCreate))
	mux.Handle("DELETE /v1/feed_follows/{feedFollowID}", apiConfig.middlewareAuth(apiConfig.handlerFeedFollowsDelete))

	mux.HandleFunc("POST /v1/users", apiConfig.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", apiConfig.middlewareAuth(apiConfig.handlerUsersGet))

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
	}

	const feedsCount = 10
	const scrapeFrequency = 60 * time.Second
	go startScraping(apiConfig.DB, feedsCount, scrapeFrequency)

	log.Printf("Serving on port: %s\n", serverPort)
	log.Fatal(server.ListenAndServe())
}
