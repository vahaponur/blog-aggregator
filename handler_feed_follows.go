package main

import (
	"blog-aggregator/internal/database"

	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	ctx := context.Background()
	feed_follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    uuid.UUID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &feed_follow)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	data, err := cfg.DB.CreateFeedFollow(ctx, feed_follow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, data)

}
func deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	ctx := context.TODO()
	feedFollowIdstr := mux.Vars(r)["feedFollowID"]
	ffId, err := uuid.Parse(feedFollowIdstr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	relExist, err := cfg.DB.FeedFollowExist(ctx, database.FeedFollowExistParams{
		UserID: user.ID,
		ID:     ffId,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	if !relExist {
		respondWithError(w, http.StatusBadRequest, errors.New("User don't follow given feed"))
		return
	}
	err = cfg.DB.DeleteFeedFollow(ctx, ffId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, "Unfollowed")

}
func getFollowedFeedsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	ctx := context.Background()
	feeds, err := cfg.DB.GetAllFeedsOfUser(ctx, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
