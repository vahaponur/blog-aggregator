package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(payload)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}
func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, err.Error())
}
