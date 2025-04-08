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

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samoei-ftw/specno/backend/common/utils"
	"github.com/samoei-ftw/specno/backend/project_service/internal/handlers"
	repo "github.com/samoei-ftw/specno/backend/project_service/internal/repository"
	"github.com/samoei-ftw/specno/backend/project_service/internal/services"
)

func main() {
	if err := utils.InitializeDatabase(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	projectRepo := repo.NewProjectRepository(utils.GetDB())

	// Initialize project service with the repository
	projectService := services.NewProjectService(projectRepo)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/projects", handlers.CreateProjectHandler(projectService)).Methods("POST")
	r.HandleFunc("/projects", handlers.GetProjectHandler(projectService)).Methods("GET")
	r.HandleFunc("/projects/{user_id}", handlers.ListProjectHandler(projectService)).Methods("GET")
	
	// Use cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // TODO: remove hardcoding
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	// Start the server
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
