// Package utils
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for api response logic
package utils

import (
	"encoding/json"
	"net/http"
)

// Standard API response format
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type GetOwnerResponse struct {
	IsOwner  bool      `json:"is_owner"`
}


func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithErrorMessage(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{
		"message": message,
	})
}

func RespondWithSuccess(w http.ResponseWriter, code int, data interface{}) {
	RespondWithJSON(w, code, data)
}