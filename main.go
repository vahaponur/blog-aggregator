package main

import "net/http"

var cfg *Config

func main() {
	cfg = createConfig()
	mainRouter := createRouter()
	http.Handle("/", mainRouter)
	http.ListenAndServe(cfg.env.PORT, nil)
}
