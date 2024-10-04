package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"ajna-rfq/internal/config"
	"ajna-rfq/internal/repo/sqlite3"
	"ajna-rfq/internal/server"
	"ajna-rfq/internal/service"
)

func main() {
	cfg, err := config.ReadConfig(getEnv("CONFIG_FILE", "config.json"))
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	repo, err := sqlite3.NewOrdersRepoSQLite3(getEnv("DATABASE_FILE", "8080"))
	if err != nil {
		log.Fatalf("failed to start repo: %v", err)
	}
	svc, err := service.NewService(cfg, repo)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	log.Println("starting server")
	err = http.ListenAndServe(":"+getEnv("PORT", "8080"), cors.Default().Handler(server.NewRouter(svc)))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
