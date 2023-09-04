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
	v1Router.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan struct{})
		go readinessHandler(w, r, ch)
		<-ch
	})
	return router
}
func readinessHandler(w http.ResponseWriter, r *http.Request, ch chan<- struct{}) {

	data := map[string]interface{}{
		"status": "ok",
	}
	respondWithJSON(w, 200, data)
	ch <- struct{}{}
}
