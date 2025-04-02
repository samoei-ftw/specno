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
	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"

	"github.com/samoei-ftw/specno/backend/gateways"
	"golang.org/x/crypto/bcrypt"
)
var userGateway = gateways.UserGatewayInit()

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
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

	project, err := services.CreateProject(request.Name, request.Description, request.UserID)
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByEmail(dto.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateToken(int(user.ID)) // TODO: add role as arg
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	response := struct {
		Token  string `json:"token"`
		Role   string `json:"role"`
		UserId int    `json:"user_id"`
	}{
		Token:  token,
		Role:   user.Role,
		UserId: int(user.ID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
