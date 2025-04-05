package repository

import (
	"errors"

	projectModels "github.com/samoei-ftw/specno/backend/project_service/internal/models"

	//commonModels "github.com/samoei-ftw/specno/backend/common/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

type ProjectRepository interface {
	Create(project *projectModels.Project) error
	GetByUserID(userID int) ([]projectModels.Project, error)
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

func (r *projectRepo) GetByUserID(userID int) ([]projectModels.Project, error) {
	if DB == nil {
		return nil, errors.New("DB connection error in fetch user.")
	}
	var projects []projectModels.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
