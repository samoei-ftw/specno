// Package services
// Author: Samoei Oloo
// Created: 2025-04-02
// License: None
//
// This script is responsible for the service layer for the project
package services

import (
	"errors"
	"log"

	"github.com/samoei-ftw/specno/backend/gateways"
	"github.com/samoei-ftw/specno/backend/project_service/internal/models"
	"github.com/samoei-ftw/specno/backend/project_service/internal/repository"
)
var (
	userGateway = gateways.UserGatewayInit()
	//projectRepo = gateways.ProjectGatewayInit()
)
type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}
// Creates a project for a user
func (s *ProjectService) CreateProject(name, description string, userId int) (*models.Project, error) {
	// Validate user
	user, err := userGateway.GetUserByID(userId)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}
	project := &models.Project{
		Name:        name,
		Description: description,
		UserID:      int(user.ID),
	}
	if err := s.repo.Create(project); err != nil {
		log.Printf("Error creating project: %v", err)
		return nil, errors.New("failed to create project")
	}

	return project, nil
}
