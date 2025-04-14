package internal

import (
	"github.com/samoei-ftw/specno/backend/common/models"
)
type Repo interface {
	Create(task *models.Task) error
	ListTasksForProject(projectId uint) ([]models.Task, error)
	GetTaskById(taskId uint) (models.Task, error)
}