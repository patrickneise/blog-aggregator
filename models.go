package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/patrickneise/blog-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		URL:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFollowToFollow(follow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        follow.ID,
		CreatedAt: follow.CreatedAt,
		UpdatedAt: follow.UpdatedAt,
		UserID:    follow.UserID,
		FeedID:    follow.FeedID,
	}
}

func databaseFollowsToFollows(follows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(follows))
	for i, follow := range follows {
		result[i] = databaseFollowToFollow(follow)
	}
	return result
}
