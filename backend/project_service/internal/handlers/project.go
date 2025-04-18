// Package handlers
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the handlers for the project api
package internal

import (
	"encoding/json"
	"log"

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
			utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		log.Printf("project request user id: %d", projectCreateRequest.UserID)

		if projectCreateRequest.Name == "" || projectCreateRequest.Description == "" || projectCreateRequest.UserID == 0 {
			utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Project name and valid user ID are required")
			return
		}

		project, err := service.CreateProject(projectCreateRequest.Name, projectCreateRequest.Description, projectCreateRequest.UserID)
		if err != nil {
			if err.Error() == "user not found" {
				utils.RespondWithErrorMessage(w, http.StatusNotFound, "User not found")
				return
			}
			if err.Error() == "no permissions"{
			utils.RespondWithErrorMessage(w, http.StatusUnauthorized, "User is not allowed to create a project. Please login with different credentials")}
			return
		}

		utils.RespondWithSuccess(w, http.StatusCreated, map[string]interface{}{
			"id":      project.ID,
			"message": "Project created successfully",
		})
	}
}

func ListProjectHandler(service *service.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := mux.Vars(r)["user_id"]
		userId, err := strconv.Atoi(userIdStr)
		if err != nil || userId <= 0 {
			utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid or missing user ID")
			return
		}

		projects, err := service.ListProjects(uint(userId))
		if err != nil {
			status := http.StatusInternalServerError
			/**if errors.Is(err, service.ErrUserNotFound) {
				status = http.StatusNotFound
			}*/ //TODO
			utils.RespondWithErrorMessage(w, status, err.Error())
			return
		}

		utils.RespondWithSuccess(w, http.StatusOK, map[string]interface{}{
			"projects": projects,
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
