package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const (
	portEnvVar = "PORT"
)

func main() {
	fmt.Println("Starting web scrapper...")
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv(portEnvVar)
	if port == "" {
		log.Fatal("No port specified in .env file")
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}).Handler)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", handleError)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	fmt.Println("Server running on port:", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
