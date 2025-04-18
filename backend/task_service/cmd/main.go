// Package main
// Author: Samoei Oloo
// Created: 2025-04-09
// License: None
//
// This script is responsible for the main execution of the task service
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samoei-ftw/specno/backend/common/utils"
	handlers "github.com/samoei-ftw/specno/backend/task_service/internal/handlers"
	repo "github.com/samoei-ftw/specno/backend/task_service/internal/repository"
	services "github.com/samoei-ftw/specno/backend/task_service/internal/services"
)

func main() {
	if err := utils.InitializeDatabase(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	taskRepo := repo.NewRepository(utils.GetDB())

	taskService := services.NewService(taskRepo)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.CreateTaskHandler(taskService)).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskHandler(taskService)).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler(taskService)).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler(taskService)).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", handlers.AssignUserHandler(taskService)).Methods("PUT")
	
	// Use cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // TODO: remove hardcoding
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	// Start the server
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
