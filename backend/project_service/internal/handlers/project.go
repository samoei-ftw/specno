// Package handlers
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the handlers for the project api
package internal

import (
	"encoding/json"
	//"log"
	"net/http"
	"strconv"

	//"strings"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/utils"
	models "github.com/samoei-ftw/specno/backend/project_service/internal/models"
	service "github.com/samoei-ftw/specno/backend/project_service/internal/service"
)


func CreateProjectHandler(service *service.ProjectService) http.HandlerFunc {
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

func ListProjectHandler(service *service.ProjectService) http.HandlerFunc {
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

func GetProjectHandler(service *service.ProjectService) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			projectIdStr := vars["id"]

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

func UpdateProjectHandler(service *service.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from url path
		vars := mux.Vars(r)
		projectIdStr := vars["id"]

		/** query param extraction
		projectIdStr := r.URL.Query().Get("id") */

		projectId, err := strconv.Atoi(projectIdStr)
		// input validation
		if err != nil || projectId <= 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid project ID",
			})
			return
		}
		var projectUpdateRequest models.ProjectUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&projectUpdateRequest); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Bad request",
			})
			return
		}

		// Call the service layer to update the project
		project, err := service.UpdateProject(projectId, projectUpdateRequest)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to update project",
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

func DeleteProjectHandler(service *service.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectIdStr := vars["id"]

		projectId, err := strconv.Atoi(projectIdStr)
		if err != nil || projectId <= 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid project ID",
			})
			return
		}
		isDeleted, err := service.DeleteProject(projectId)
		if isDeleted != true {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to delete project",
			})
			return
		}

	utils.RespondWithJSON(w, http.StatusOK, utils.Response{
		Status: "success",
		Data:   "Project deleted", // TODO: fix this
	})
}
}
