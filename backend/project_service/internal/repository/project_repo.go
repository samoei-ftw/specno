package repository

import (
	"errors"

	projectModels "github.com/samoei-ftw/specno/backend/project_service/internal/models"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *projectModels.Project) error
	GetByUserID(userID uint) ([]projectModels.Project, error)
	GetProjectById(projectId uint) (projectModels.Project, error)
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(project *projectModels.Project) error {
	if r.db == nil {
		return errors.New("DB connection error in create.")
	}
	return r.db.Create(project).Error
}

func (r *projectRepo) GetByUserID(userID uint) ([]projectModels.Project, error) {
	if r.db == nil {
		return nil, errors.New("DB connection error in fetch user.")
	}

	var projects []projectModels.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
func (r *projectRepo) GetProjectById(projectId uint) (projectModels.Project, error) {
	if r.db == nil {
		return projectModels.Project{}, errors.New("DB connection error")
	}

	var project projectModels.Project
	if err := r.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		return projectModels.Project{}, err
	}

	return project, nil
}