// Package handlers
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the handlers for the project api
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samoei-ftw/specno/backend/project_service/internal/services"

	"github.com/samoei-ftw/specno/backend/common/utils"

	"github.com/samoei-ftw/specno/backend/gateways"
)
var userGateway = gateways.UserGatewayInit()

func CreateProjectHandler(service *services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			UserID      int    `json:"user_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid request payload",
			})
			return
		}

		if request.Name == "" || request.Description == "" || request.UserID == 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Project name and valid user ID are required",
			})
			return
		}

		project, err := service.CreateProject(request.Name, request.Description, request.UserID)
		if err != nil {
			if err.Error() == "user not found" {
				utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
					Status:  "error",
					Message: "User not found",
				})
				return
			}
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to create project",
			})
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, utils.Response{
			Status: "success",
			Data:   project,
		})
	}
}

