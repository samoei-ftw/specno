package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	"tasko/internal/handlers"

	//"tasko/internal/models"

	"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/register", handlers.RegisterHandler)
	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
