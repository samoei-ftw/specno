package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"tasko/internal/models"
	"tasko/internal/user"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	models.InitDB()

	// Define routes
	http.HandleFunc("/register", user.RegisterHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
