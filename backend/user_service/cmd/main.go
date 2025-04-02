package main

import (
	"log"
	"net/http"
	"os"

	"github.com/samoei-ftw/specno/backend/user_service/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	err := godotenv.Load("/backend/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/users/{arg:[0-9]+}", handlers.FetchUserHandler)

	// Use cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // TODO: remove hardcoding
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(mux)

	// Start the server
	log.Println("Starting server on port", port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
