package main

import (
	"fmt"
	"net/http"
)

var cfg *Config

func main() {
	cfg = createConfig()
	mainRouter := createRouter()
	http.Handle("/", mainRouter)
	fmt.Println(fmt.Sprintf("Server started at port:%v", cfg.env.PORT))
	http.ListenAndServe(cfg.env.PORT, nil)
}
