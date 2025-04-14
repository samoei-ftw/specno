package internal

import (
	"errors"

	common "github.com/samoei-ftw/specno/backend/common/models"
	interfaces "github.com/samoei-ftw/specno/backend/project_service/internal/interfaces"

	//models "github.com/samoei-ftw/specno/backend/project_service/internal/models"
	"gorm.io/gorm"
)
type projectRepo struct {
	db *gorm.DB
}
func NewProjectRepository(db *gorm.DB) interfaces.Repository {
    return &projectRepo{db: db}
}

func (r *projectRepo) Create(project *common.Project) error {
	if r.db == nil {
		return errors.New("DB connection error in create.")
	}
	return r.db.Create(project).Error
}

func (r *projectRepo) GetByUserID(userID uint) ([]common.Project, error) {
	if r.db == nil {
		return nil, errors.New("DB connection error in fetch user.")
	}

	var projects []common.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
func (r *projectRepo) GetProjectById(projectId uint) (common.Project, error) {
	if r.db == nil {
		return common.Project{}, errors.New("DB connection error")
	}

	var project common.Project
	if err := r.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		return common.Project{}, err
	}

	return project, nil
}
func (r *projectRepo) GetUserForProject(projectId uint) (uint, error) {
	if r.db == nil {
		return 0, errors.New("DB connection error")
	}

	var project common.Project
	if err := r.db.Where("id = ?", projectId).First(&project).Error; err != nil {
		return 0, err
	}

	return project.UserID, nil
}

func (r *projectRepo) UpdateProject(project *common.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepo) DeleteProject(project *common.Project) (bool, error) {
	err := r.db.Delete(project).Error
	if err != nil {
		return false, err
	}
	return true, nil
}