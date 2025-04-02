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

	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"

	"golang.org/x/crypto/bcrypt"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      int    `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload.", http.StatusBadRequest)
		return
	}

	if dto.Name == "" {
		http.Error(w, "Project name is required.", http.StatusBadRequest)
		return
	}

	userID, err := services.CreateProject(dto.Email, dto.Password)
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      userID,
		"message": "User registered successfully",
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
