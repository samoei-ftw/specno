package internal

import (
	"errors"
	"log"

	common "github.com/samoei-ftw/specno/backend/common/models"
	"github.com/samoei-ftw/specno/backend/gateways"
	interfaces "github.com/samoei-ftw/specno/backend/project_service/internal/interfaces"
	models "github.com/samoei-ftw/specno/backend/project_service/internal/models"
)
type ProjectService struct {
	repo interfaces.Repository
}
var userGateway = gateways.UserGatewayInit()
func NewProjectService(repo interfaces.Repository) *ProjectService {
	return &ProjectService{repo: repo}
}
// Creates a project for a user
func (s *ProjectService) CreateProject(name, description string, userId uint) (*common.Project, error) {
	// Validate user
	user, err := userGateway.GetUserByID(userId)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	project := &common.Project{
		Name:        name,
		Description: description,
		UserID:      user.ID,
	}

	if err := s.repo.Create(project); err != nil {
		log.Printf("Error creating project: %v", err)
		return nil, errors.New("failed to create project")
	}

	return project, nil
}

// Lists all projects for a user
func (s *ProjectService) ListProjects(userId uint) ([]*common.Project, error) {
	// Validate user exists
	user, err := userGateway.GetUserByID(userId)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	projectList, err := s.repo.GetByUserID(user.ID)
	if err != nil {
		log.Printf("Error fetching projects for user %d: %v", userId, err)
		return nil, errors.New("failed to retrieve projects")
	}

	projects := make([]*common.Project, len(projectList))
	for i, p := range projectList {
		projects[i] = &p
	}

	return projects, nil
}

func (s *ProjectService) GetProject(projectId uint) (*common.Project, error) {
	project, err := s.repo.GetProjectById(projectId)
	if err != nil {
		log.Printf("Error fetching project. %d: %v", projectId, err)
		return nil, errors.New("failed to retrieve project")
	}
	return &project, nil
}

func (s *ProjectService) GetUserForProject(projectId uint, bearer string) (uint, error) {
	userId, err := s.repo.GetUserForProject(projectId)
	if err != nil {
		log.Printf("Error fetching user for project %d: %v", projectId, err)
		return 0, errors.New("failed to retrieve user for project")
	}
	return userId, nil
}

// UpdateProject updates the project with the given ID
func (s *ProjectService) UpdateProject(projectId int, request models.ProjectUpdateRequest) (*common.Project, error) {
	project, err := s.repo.GetProjectById(uint(projectId))
	if err != nil {
		return nil, err
	}
	project.Name = request.Name
	project.Description = request.Description

	err = s.repo.UpdateProject(&project)
	if err != nil {
		return nil, errors.New("Failed to update project")
	}

	return &project, nil
}

// Deletes a project with the given id
func (s *ProjectService) DeleteProject(projectId int) (bool, error) {
	project, err := s.repo.GetProjectById(uint(projectId))
	if err != nil {
		return false, err
	}

	isDeleted, err := s.repo.DeleteProject(&project)
	if err != nil {
		return false, errors.New("Failed to delete project")
	}

	return isDeleted, nil
}