// Package handlers
// Author: Samoei Oloo
// Created: 2025-04-09
// License: None
//
// This script is responsible for the handlers for the task api
package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/samoei-ftw/specno/backend/common/utils"

	"github.com/samoei-ftw/specno/backend/gateways"
)
//var userGateway = gateways.UserGatewayInit()

func CreateTaskHandler(service *service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	taskCreateRequest taskCreateRequest;
		if err := json.NewDecoder(r.Body).Decode(&taskCreateRequest); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid request payload",
			})
			return
		}

		if taskCreateRequest.Title == "" || taskCreateRequest.Description == "" || taskCreateRequest.ProjectID == 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Validation failure",
			})
			return
		}

		task, err := service.CreateTask(projectCreateRequest.Name, projectCreateRequest.Description, projectCreateRequest.UserID)
		if err != nil {
			if err.Error() == "user not found" {
				utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
					Status:  "error",
					Message: "Not found error",
				})
				return
			}
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to create task",
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

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid user ID",
			})
			return
		}

		projects, err := service.ListProjects(uint(userId))
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

	func GetProjectHandler(service *services.ProjectService) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			projectIdStr := r.URL.Query().Get("id")

			projectId, err := strconv.Atoi(projectIdStr)
			if err != nil || projectId <= 0 {
				utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
					Status:  "error",
					Message: "Invalid project ID",
				})
				return
			}

			project, err := service.GetProject(uint(projectId))
			if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to retrieve project",
			})
			return
		}
		if project == nil {
			utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
				Status:  "error",
				Message: "Project not found",
			})
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data:   project,
		})
	}
}
