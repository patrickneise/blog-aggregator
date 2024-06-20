package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickneise/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type response struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode parameters")
		return
	}

	now := time.Now().UTC()
	feedArgs := &database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	}

	feed, err := cfg.DB.CreateFeed(req.Context(), *feedArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	followArgs := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := cfg.DB.CreateFeedFollow(req.Context(), *followArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed follow")
		return
	}

	respondWithJSON(w, http.StatusAccepted, response{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseFollowToFollow(follow)},
	)
}

func (cfg *apiConfig) handlerFeedsGetAll(w http.ResponseWriter, req *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retreive feeds")
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

// func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, req *http.Request, user database.User) {
// 	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
// }
