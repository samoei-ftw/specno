// Package handlers
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the handlers for the project api
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samoei-ftw/specno/backend/project_service/internal/services"

	"github.com/samoei-ftw/specno/backend/gateways"
)
var userGateway = gateways.UserGatewayInit()

func CreateProjectHandler(service *services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			UserID      int    `json:"user_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Printf("Failed to decode payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Decoded payload: %+v", payload)

		project, err := service.CreateProject(payload.Name, payload.Description, payload.UserID, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(project)
	}
}

