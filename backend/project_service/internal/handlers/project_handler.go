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
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/project_service/internal/services"

	"github.com/samoei-ftw/specno/backend/common/utils"

	"github.com/samoei-ftw/specno/backend/gateways"
)
var userGateway = gateways.UserGatewayInit()

func CreateProjectHandler(service *services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var projectCreateRequest struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			UserID      uint    `json:"user_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&projectCreateRequest); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid request payload",
			})
			return
		}

		if projectCreateRequest.Name == "" || projectCreateRequest.Description == "" || projectCreateRequest.UserID == 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Project name and valid user ID are required",
			})
			return
		}

		project, err := service.CreateProject(projectCreateRequest.Name, projectCreateRequest.Description, projectCreateRequest.UserID)
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

func ListProjectHandler(service *services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r) 
		userIdStr := vars["user_id"] 

		// Check if user_id is provided
		if userIdStr == "" {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Missing user_id path parameter",
			})
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid user_id format",
			})
			return
		}

		// Call the service to fetch the projects
		projects, err := service.ListProjects(uint(userId))
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to retrieve projects",
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data:   projects,
		})
	}
}
