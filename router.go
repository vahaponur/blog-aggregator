package main

import (
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	//Get cors option from config
	router.Use(cors.Handler(cfg.corsOpt))
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/route1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is route 1 of the subrouter"))
	})
	return router
}
