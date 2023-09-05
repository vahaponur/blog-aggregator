package main

import (
	"blog-aggregator/internal/database"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

func createFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	feed := database.CreateFeedParams{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	feed.ID = uuid.New()
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = time.Now()
	feed.UserID = user.ID
	ctx := context.Background()
	data, err := cfg.DB.CreateFeed(ctx, feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	data2, err := cfg.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    data.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	type CreateFeedResponse struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}
	responseModel := CreateFeedResponse{
		Feed:       data,
		FeedFollow: data2,
	}

	respondWithJSON(w, http.StatusOK, responseModel)
}
func getAllFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	feeds, err := cfg.DB.GetAllFeeds(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
