package main

import (
	"errors"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	//Get cors option from config
	router.Use(cors.Handler(cfg.corsOpt))
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan struct{})
		go readinessHandler(w, r, ch)
		<-ch
	}).Methods("GET")
	v1Router.HandleFunc("/err", errorHandler)
	userRouter := v1Router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", createUserHandler).Methods("POST")
	userRouter.HandleFunc("", cfg.middlewareAuth(getUserByApiKey)).Methods("GET")
	feedRouter := v1Router.PathPrefix("/feeds").Subrouter()
	feedRouter.HandleFunc("", cfg.middlewareAuth(createFeed)).Methods("POST")
	feedRouter.HandleFunc("", getAllFeeds).Methods("GET")
	feed_followRouter := v1Router.PathPrefix("/feed_follows").Subrouter()
	feed_followRouter.HandleFunc("", cfg.middlewareAuth(createFeedFollow)).Methods("POST")
	feed_followRouter.HandleFunc("/{feedFollowID}", cfg.middlewareAuth(deleteFeedFollow)).Methods("DELETE")
	feed_followRouter.HandleFunc("", cfg.middlewareAuth(getFollowedFeedsByUser)).Methods("GET")
	postRouter := v1Router.PathPrefix("/posts").Subrouter()
	postRouter.HandleFunc("", cfg.middlewareAuth(getPostsByUser)).Methods("GET")
	return router
}

// Test responseWithJson
func readinessHandler(w http.ResponseWriter, r *http.Request, ch chan<- struct{}) {

	data := map[string]interface{}{
		"status": "ok",
	}
	respondWithJSON(w, http.StatusOK, data)
	ch <- struct{}{}
}
func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, errors.New("Internal Server Error"))
}
