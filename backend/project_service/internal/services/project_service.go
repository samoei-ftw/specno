// Package services
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the service layer for the project
package services

import (
	"errors"

	"github.com/samoei-ftw/specno/backend/project_service/internal/models"
	"github.com/samoei-ftw/specno/backend/user_service/internal/repo"
)

// Creates a project for a user
func CreateProject(name, description string, userId int) (*models.Project, error) {
	// Validate user
	user, err := repo.GetUserByID(userId)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
