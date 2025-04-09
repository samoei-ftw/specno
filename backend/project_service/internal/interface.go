package internal

import (
	"github.com/samoei-ftw/specno/backend/common/models"
)
type Repository interface {
	Create(project *models.Project) error
	GetByUserID(userID uint) ([]models.Project, error)
	GetProjectById(projectId uint) (models.Project, error)
	GetUserForProject(projectId uint) (uint, error)
}