package internal

import (
	"errors"
	"fmt"

	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/gorm"
)
type Repo struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) Repo {
	return Repo{db: db}
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

func (r *Repo) UpdateTask(
	taskId uint,
	status *enums.TaskStatus,
	title *string,
	description *string,
	assignee *uint,
) (models.Task, error) {
	if r.db == nil {
		return models.Task{}, errors.New("DB connection error")
	}

	var task models.Task
	if err := r.db.Where("id = ?", taskId).First(&task).Error; err != nil {
		return models.Task{}, err
	}
// only update fields provided - find better way to do this
	if status != nil {
		task.Status = status.String()
	}
	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if assignee != nil {
		task.UserID = *assignee
	}

	if err := r.db.Save(&task).Error; err != nil {
		fmt.Printf("Failed to save task with ID %d: %v\n", taskId, err)
		return models.Task{}, err
	}
	return task, nil
}

func (r *Repo) DeleteTask(task *models.Task) (bool, error) {
	err := r.db.Delete(task).Error
	if err != nil {
		return false, err
	}
	return true, nil
}