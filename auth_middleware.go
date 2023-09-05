package main

import (
	"blog-aggregator/internal/database"
	"context"
	"net/http"
	"strings"
)

type authMW func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *Config) middlewareAuth(handler authMW) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyStr := r.Header.Get("Authorization")

		apikey, ok := strings.CutPrefix(keyStr, "ApiKey ")
		if !ok {
			respondWithError(w, http.StatusUnauthorized, YouShallNotPass())
			return
		}
		ctx := context.Background()
		user, err := cfg.DB.GetUserByApikey(ctx, apikey)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				respondWithError(w, http.StatusNotFound, err)
				return
			}
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		handler(w, r, user)

	}
}
