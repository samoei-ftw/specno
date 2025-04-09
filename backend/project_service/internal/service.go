package internal

import (
	"errors"
	"log"

	"github.com/samoei-ftw/specno/backend/common/models"
	"github.com/samoei-ftw/specno/backend/gateways"
)

var userGateway = gateways.UserGatewayInit()
func NewProjectService(repo Repository) *ProjectService {
	return &ProjectService{repo: repo}
}
// Creates a project for a user
func (s *ProjectService) CreateProject(name, description string, userId uint) (*models.Project, error) {
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
		UserID:      user.ID,
	}

	if err := s.repo.Create(project); err != nil {
		log.Printf("Error creating project: %v", err)
		return nil, errors.New("failed to create project")
	}

	return project, nil
}

// Lists all projects for a user
func (s *ProjectService) ListProjects(userId uint) ([]*models.Project, error) {
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

	projects := make([]*models.Project, len(projectList))
	for i, p := range projectList {
		projects[i] = &p
	}

	return projects, nil
}

func (s *ProjectService) GetProject(projectId uint) (*models.Project, error) {
	project, err := s.repo.GetProjectById(projectId)
	if err != nil {
		log.Printf("Error fetching project. %d: %v", projectId, err)
		return nil, errors.New("failed to retrieve project")
	}
	return &project, nil
}

func (s *ProjectService) GetUserForProject(projectId uint) (uint, error) {
	userId, err := s.repo.GetUserForProject(projectId)
	if err != nil {
		log.Printf("Error fetching user for project %d: %v", projectId, err)
		return 0, errors.New("failed to retrieve user for project")
	}
	return userId, nil
}