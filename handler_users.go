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

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	param, err := parseCreateUserParams(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("Bad request"))
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	type Response struct {
		user database.User
		err  error
	}
	resChan := make(chan Response)

	go func() {
		user, err := cfg.DB.CreateUser(ctx, param)
		resChan <- Response{err: err, user: user}
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
			respondWithJSON(w, http.StatusCreated, dbUserToUser(resp.user))
			return
		}

	}

}
func parseCreateUserParams(r *http.Request) (database.CreateUserParams, error) {
	param := database.CreateUserParams{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	err = json.Unmarshal(body, &param)
	if err != nil {
		return param, err
	}

	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	param.ID = uuid.New()

	return param, nil
}
func getUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, dbUserToUser(user))
}
