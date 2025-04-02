package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/samoei-ftw/specno/backend/user_service/internal/handlers"
	"github.com/samoei-ftw/specno/backend/user_service/internal/repo"

	"github.com/rs/cors"
)

func main() {
	// Load environment variables - TODO
	jwtSecret := os.Getenv("JWT_SECRET")
	if err := repo.InitializeDatabase(); err != nil {
		log.Fatal("Database initialization failed:", err)
	}

	// Create the UserRepository with the initialized DB
	userRepository := repo.NewUserRepository(repo.GetDB())


	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
    r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/users/{id:[0-9]+}", handlers.FetchUserHandler).Methods("GET")
	
	// Use cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // TODO: remove hardcoding
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(r)

	// Start the server
	log.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
