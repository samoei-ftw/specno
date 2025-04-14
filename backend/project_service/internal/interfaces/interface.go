package interfaces

import (
	common "github.com/samoei-ftw/specno/backend/common/models"
)

type Repository interface {
	Create(project *common.Project) error
	GetByUserID(userID uint) ([]common.Project, error)
	GetProjectById(projectId uint) (common.Project, error)
	GetUserForProject(projectId uint) (uint, error)
	UpdateProject(project *common.Project) error
	DeleteProject(project *common.Project) (bool, error)
}