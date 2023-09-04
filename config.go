package main

import (
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	env     Environment
	corsOpt cors.Options
}
type Environment struct {
	PORT string
}

func createConfig() *Config {
	cfg := &Config{}

	env, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	cfg.env = env
	cfg.corsOpt = cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}
	return cfg
}

// Loads environment variables from .env file
func loadEnv() (env Environment, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	env = Environment{}
	env.PORT = os.Getenv("PORT")
	return env, nil
}
