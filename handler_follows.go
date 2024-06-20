package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickneise/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode parameters")
		return
	}

	now := time.Now().UTC()
	args := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    params.FeedID,
	}

	follow, err := cfg.DB.CreateFeedFollow(req.Context(), *args)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed follow")
		return
	}

	respondWithJSON(w, http.StatusAccepted, databaseFollowToFollow(follow))
}

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, req *http.Request, user database.User) {
	follows, err := cfg.DB.GetUserFeedFollows(req.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retreive feeds")
	}

	respondWithJSON(w, http.StatusOK, databaseFollowsToFollows(follows))
}

func (cfg *apiConfig) handlerFeedFollowsDelete(w http.ResponseWriter, req *http.Request, user database.User) {
	followID, err := uuid.Parse(req.PathValue("feedFollowID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse feed follow id")
	}

	err = cfg.DB.DeleteFeedFollow(req.Context(), followID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete follow")
	}

	w.WriteHeader(http.StatusNoContent)
}
