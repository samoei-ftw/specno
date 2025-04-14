package internal

import (
	"errors"

	"github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/gorm"
)
type Repo struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) Repository {
	return &Repo{db: db}
}

func (r *Repo) Create(task *models.Task) error {
	if r.db == nil {
		return errors.New("DB connection not initialized")
	}
	return r.db.Create(task).Error
}

func (r *Repo) ListTasksForProject(projectId uint) ([]models.Task, error) {
	if r.db == nil {
		return nil, errors.New("DB connection error")
	}

	var tasks []models.Task
	if err := r.db.Where("project_id = ?", projectId).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repo) GetTaskById(taskId uint) (models.Task, error) {
	if r.db == nil {
		return models.Task{}, errors.New("DB connection error")
	}

	var task models.Task
	if err := r.db.Where("id = ?", taskId).First(&task).Error; err != nil {
		return models.Task{}, err
	}

	return task, nil
}