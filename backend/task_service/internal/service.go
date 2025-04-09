package internal

import (
	"errors"
	"log"
	"time"

	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/models"
	"github.com/samoei-ftw/specno/backend/gateways"
)

var (
	userGateway = gateways.UserGatewayInit()
	projectGateway = gateways.ProjectGatewayInit()
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

// Adds a task to a project
func (s *service) CreateTask(title, description string, userId uint, projectId uint, status enums.TaskStatus) (*models.Task, error) {
	//Validate project belongs to user
	project, err := projectGateway.UserOwnsProject(projectId)
	if err != nil {
		log.Printf("Gateway error %d: %v", projectId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("Gateway error")
		}
		return nil, errors.New("Gateway error")
	}

	task := &models.Task{
		Title:        title,
		Description: description,
		UserID:      project.UserID,
		ProjectID: project.ID,
		CreatedAt: time.Now(),
		Status: status,
	}

	if err := s.repo.Create(task); err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, errors.New("failed to create task")
	}

	return task, nil
}

// Lists all tasks for a project
func (s *service) ListTasks(projectId uint) ([]*models.Task, error) {
	taskList, err := s.repo.ListTasksForProject(projectId)
	if err != nil {
		log.Printf("Error fetching tasks for project %d: %v", projectId, err)
		return nil, errors.New("failed to retrieve tasks")
	}
	tasks := make([]*models.Task, len(taskList))
	for i, t := range taskList {
		tasks[i] = &t
	}

	return tasks, nil
}

func (s *service) GetTask(taskId uint) (*models.Task, error){
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		log.Printf("Error fetching task. %d: %v", taskId, err)
		return nil, errors.New("failed to retrieve task")
	}
	return &task, nil
}