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

func (g *UserGateway) ValidateUserId(userID uint) (bool, error) {
	url := fmt.Sprintf("%s/users/%d", g.BaseURL, userID)

	resp, err := g.HTTPClient.Get(url)
	if err != nil {
		return false, fmt.Errorf("Gateway error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("User service returned status: %d", resp.StatusCode)
	}

	var userResponse models.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return false, fmt.Errorf("Error interpreting response: %w", err)
	}
	return true, nil
}
