package internal

import (
	"errors"

	models "github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/gorm"
)

func NewProjectRepository(db *gorm.DB) Repository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(project *models.Project) error {
	if r.db == nil {
		return errors.New("DB connection error in create.")
	}
	return r.db.Create(project).Error
}

func (r *projectRepo) GetByUserID(userID uint) ([]models.Project, error) {
	if r.db == nil {
		return nil, errors.New("DB connection error in fetch user.")
	}

	var projects []models.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
func (r *projectRepo) GetProjectById(projectId uint) (models.Project, error) {
	if r.db == nil {
		return models.Project{}, errors.New("DB connection error")
	}

	var project models.Project
	if err := r.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		return models.Project{}, err
	}

	return project, nil
}
func (r *projectRepo) GetUserForProject(projectId uint) (uint, error) {
	if r.db == nil {
		return 0, errors.New("DB connection error")
	}

	var project models.Project
	if err := r.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		return 0, err
	}

	return project.UserID, nil
}