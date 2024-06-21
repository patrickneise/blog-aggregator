package main

import (
	"net/http"
	"strconv"

	"github.com/patrickneise/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, req *http.Request, user database.User) {
	limitStr := req.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	args := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := cfg.DB.GetPostsByUser(req.Context(), args)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retreiving posts")
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
