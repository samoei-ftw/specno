package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/samoei-ftw/specno/backend/user_service/internal/handlers"
	"github.com/samoei-ftw/specno/backend/user_service/internal/repo"
	"github.com/samoei-ftw/specno/backend/user_service/internal/services"

	"github.com/rs/cors"
)

func main() {
	if err := repo.InitializeDatabase(); err != nil {
		log.Fatal("Database initialization failed:", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	userRepo := repo.NewUserRepository(repo.GetDB())
    userService := services.NewUserService(userRepo)
	r := mux.NewRouter()
    // Register routes with handlers, injecting the userService instance
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, userService)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userService)
	}).Methods("POST")

	r.HandleFunc("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FetchUserHandler(w, r, userService)
	}).Methods("GET")
	
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
