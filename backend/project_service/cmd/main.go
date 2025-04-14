// Package main
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the main execution of the project service
package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samoei-ftw/specno/backend/common/utils"
	internal "github.com/samoei-ftw/specno/backend/project_service/internal/handlers"
	repo "github.com/samoei-ftw/specno/backend/project_service/internal/repository"
	service "github.com/samoei-ftw/specno/backend/project_service/internal/service"
	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"
)

func main() {
	if err := utils.InitializeDatabase(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		
		log.Fatal("Error loading port number from environment variables")
	}
	projectRepo := repo.NewProjectRepository(utils.GetDB())

	// Initialize project service with the repository
	projectService := service.NewProjectService(projectRepo)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/projects", internal.CreateProjectHandler(projectService)).Methods("POST")
	r.HandleFunc("/projects/{id}", internal.GetProjectHandler(projectService)).Methods("GET")
	r.HandleFunc("/projects/{id}", internal.UpdateProjectHandler(projectService)).Methods("PUT")
	r.HandleFunc("/projects/{id}", internal.DeleteProjectHandler(projectService)).Methods("DELETE")

	r.HandleFunc("/projects/{user_id}", internal.ListProjectHandler(projectService)).Methods("GET") // TODO: rename for user
	r.Handle("/projects/{project_id}/ownership", auth.JWTMiddleware(internal.GetProjectOwnerHandler(projectService))).Methods("GET") // TODO: remove
	
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
