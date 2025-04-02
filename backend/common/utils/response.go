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

// Return formatted JSON response
func RespondWithJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}