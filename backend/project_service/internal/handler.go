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
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/utils"
	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"
)


func CreateProjectHandler(service *ProjectService) http.HandlerFunc {
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

func ListProjectHandler(service *ProjectService) http.HandlerFunc {
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

func GetProjectHandler(service *ProjectService) http.HandlerFunc {
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

func GetProjectOwnerHandler(service *ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		vars := mux.Vars(r)
		projectIdStr := vars["project_id"]
		//var rf := auth.ExtractTokenFromHeader(r);
		projectId, err := strconv.Atoi(projectIdStr)
		log.Printf("Invalid project id, %s", projectIdStr)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid project ID:",
			})
			return
		}
		projectUserId, err := service.GetUserForProject(uint(projectId), tokenStr) 
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Error fetching user for project",
			})
			return
		}
		claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
		if !ok || claims == nil {
			utils.RespondWithJSON(w, http.StatusUnauthorized, utils.Response{
				Status:  "error",
				Message: "Unauthorized: user ID not found",
			})
			return
		}
		log.Printf("Claims User ID: %d, Project Owner ID: %d", claims.UserID, projectUserId)
		if uint(claims.UserID) != projectUserId {
			utils.RespondWithJSON(w, http.StatusForbidden, utils.Response{
				Status:  "error",
				Message: "User is not the project owner",
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data: utils.GetOwnerResponse{
				IsOwner: true,
			},
		})
	}
}
