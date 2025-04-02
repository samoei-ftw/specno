// Package main
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the main execution of the project service
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/samoei-ftw/specno/backend/project_service/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
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
		port = "8081"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/project", handlers.CreateProjectHandler)

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
