package main

import (
	"blog-aggregator/internal/database"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	param := database.CreateUserParams{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	json.Unmarshal(body, &param)

	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	param.ID = uuid.New()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	type UserJSON struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
	}
	type Response struct {
		user UserJSON
		err  error
	}
	resChan := make(chan Response)

	go func() {
		user, err := cfg.DB.CreateUser(ctx, param)
		resChan <- Response{err: err, user: UserJSON{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
		}}
	}()

	for {
		select {
		case <-ctx.Done():
			respondWithError(w, http.StatusGatewayTimeout, errors.New("Database took too long to response"))
			return
		case resp := <-resChan:
			if resp.err != nil {
				respondWithError(w, http.StatusInternalServerError, errors.New("Something went wrong"))
				return
			}
			respondWithJSON(w, http.StatusCreated, resp.user)
			return
		}

	}

}
