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

type ProjectGateway struct {
	BaseURL    string
	HTTPClient *http.Client
}

func ProjectGatewayInit() *ProjectGateway {
	baseURL := os.Getenv("PROJECT_SERVICE_BASE_URL")
	if baseURL == "" {
		baseURL = "http://project-service:8082"
	}

	return &ProjectGateway{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (g *ProjectGateway) UserOwnsProject(projectId uint) (*models.Project, error) {
	url := fmt.Sprintf("%s/projects/%d/ownership", g.BaseURL, projectId)

	resp, err := g.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Gateway error: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Project service returned status: %d", resp.StatusCode)
	}
	
	var ownershipResponse struct {
		IsOwner bool `json:"isOwner"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&ownershipResponse); err != nil {
		return nil, fmt.Errorf("Error interpreting response: %w", err)
	}
	if !ownershipResponse.IsOwner {
		return nil, fmt.Errorf("User does not have access to this project")
	}
	
	return nil, nil
}
