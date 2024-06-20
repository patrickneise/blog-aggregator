package main

import (
	"net/http"

	"github.com/patrickneise/blog-aggregator/internal/auth"
	"github.com/patrickneise/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetAPIKey(req.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "No API Key")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		}

		handler(w, req, user)
	}
}
