// Package gateways
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script provides methods to communicate with the User Service.

package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/samoei-ftw/specno/backend/common/models"
)

type UserGateway struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Initializes a new UserGateway
func UserGatewayInit() *UserGateway {
	baseURL := os.Getenv("USER_SERVICE_BASE_URL")
	if baseURL == "" {
		baseURL = "http://user-service:8080"
	}

	return &UserGateway{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// GetUserByID fetches a user by ID from the User Service.
func (g *UserGateway) GetUserByID(userID int, incomingRequest *http.Request) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%d", g.BaseURL, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	// Propagate the Authorization header (Bearer token)
	if authHeader := incomingRequest.Header.Get("Authorization"); authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Gateway error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("User service returned status: %d", resp.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("Error interpreting response: %w", err)
	}

	return &user, nil
}
