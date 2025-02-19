package main

import (
	"github.com/joho/godotenv"
	"log"
	"test-task/internal/app/http"
	"test-task/internal/config"
	"test-task/pkg/db/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()
	pgClient, err := postgres.NewPgConnection(&cfg.Database)
	if err != nil {
		panic(err)
	}

	http.NewApp(pgClient, cfg).Start()
}
