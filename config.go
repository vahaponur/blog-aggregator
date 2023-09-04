package main

import (
	"blog-aggregator/internal/database"
	"database/sql"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	env     Environment
	corsOpt cors.Options
	DB      *database.Queries
}
type Environment struct {
	PORT    string
	DB_CONN string
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
	db, err := sql.Open("postgres", cfg.env.DB_CONN)
	if err != nil {
		log.Fatal(err)
	}
	cfg.DB = database.New(db)
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
	env.DB_CONN = os.Getenv("DB_CONN")
	return env, nil
}
