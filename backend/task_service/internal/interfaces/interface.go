package internal

import (
	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/models"
)

type Repository interface {
	Create(task *models.Task) error
	ListTasksForProject(projectId uint) ([]models.Task, error)
	GetTaskById(taskId uint) (models.Task, error)
	UpdateTaskStatus(taskId uint, status enums.TaskStatus) (models.Task, error)
	DeleteProject(task *models.Task) (bool, error)
}

