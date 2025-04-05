package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/samoei-ftw/specno/backend/common/pkg/auth"
	utils "github.com/samoei-ftw/specno/backend/common/utils"
	"github.com/samoei-ftw/specno/backend/user_service/internal/handlers"
	"github.com/samoei-ftw/specno/backend/user_service/internal/repo"
	"github.com/samoei-ftw/specno/backend/user_service/internal/services"
)

func main() {
	if err := utils.InitializeDatabase(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err := utils.RunMigrations(os.Getenv("DB_MIGRATIONS_DIR")); err != nil {
		log.Fatal("Migrations failed:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Dependency injection
	userRepository := repo.NewUserRepository(utils.GetDB())
	userService := services.NewUserService(userRepository)

	// Router setup
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, userService)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userService)
	}).Methods("POST")

	// Protected routes
	r.Handle("/users/{id:[0-9]+}", auth.JwtAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.FetchUserHandler(w, r, userService)
	}))).Methods("GET")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // TODO: use env var in prod
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}