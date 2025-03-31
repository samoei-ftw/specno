package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"tasko/internal/handlers"
	"tasko/internal/models"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	db := models.InitDB()

	r := mux.NewRouter()

	//r.HandleFunc("/register", handlers.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
