package main

import (
	"blog-aggregator/internal/database"
	"context"
	"net/http"
)

func getPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	ctx := context.TODO()
	posts, err := cfg.DB.GetPostsByUser(ctx, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)

}
