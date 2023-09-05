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
	json.Unmarshal(body, &feed)
	feed.ID = uuid.New()
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = time.Now()
	feed.UserID = user.ID
	data, err := cfg.DB.CreateFeed(context.Background(), feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, data)
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
