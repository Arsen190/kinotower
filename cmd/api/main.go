package main

import (
	"log"
	"net/http"
	"os"

	"kinotower/internal/handler"
	"kinotower/internal/repository"
	"kinotower/internal/service"
	"kinotower/pkg/postgres"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Загружаем .env

	db, err := postgres.New(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	log.Println("SERVER STARTED ON :8080")
	http.ListenAndServe(":8080", handlers.InitRoutes())
}