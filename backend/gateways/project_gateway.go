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
	"log"
	"net/http"
	"os"
	"time"
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

func (g *ProjectGateway) UserOwnsProject(projectId uint, bearer string) (bool, error) {
	url := fmt.Sprintf("%s/projects/%d/ownership", g.BaseURL, projectId)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", bearer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", bearer)
	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("Gateway error: %w", err)
	}
	defer resp.Body.Close()
	log.Printf("Response status: %d", resp.StatusCode)
	
	if resp.StatusCode != http.StatusOK {
		return false, resp.Request.Context().Err()
	}
	
	var gatewayResponse struct {
		Status string `json:"status"`
		Data   struct {
			IsOwner bool `json:"is_owner"`
		} `json:"data"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&gatewayResponse); err != nil {
		return false, resp.Request.Context().Err()
	}
	
	isOwner := gatewayResponse.Data.IsOwner
	
	return isOwner, nil
}
