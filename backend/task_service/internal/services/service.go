package internal

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/models"
	"github.com/samoei-ftw/specno/backend/gateways"
	repository "github.com/samoei-ftw/specno/backend/task_service/internal/repository"
)

var (
	userGateway = gateways.UserGatewayInit()
	projectGateway = gateways.ProjectGatewayInit()
)

type Service struct {
	repo repository.Repo
}

func NewService(repo repository.Repo) *Service {
	return &Service{repo: repo}
}

// Adds a task to a project
func (s *Service) CreateTask(title, description string, userId uint, projectId uint, status enums.TaskStatus, bearer string) (*models.Task, error) {
	log.Printf("Validating task request project id: %s", strconv.FormatUint(uint64(projectId), 10))
	//Validate project belongs to user
	isOwner, err := projectGateway.UserOwnsProject(projectId, bearer)
	if(isOwner != true){
		return nil, errors.New("Error adding task to project. The user does not own this project.")
	}
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
		UserID:      userId,
		ProjectID: projectId,
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
func (s *Service) ListTasks(projectId uint) ([]*models.Task, error) {
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

func (s *Service) GetTask(taskId uint) (*models.Task, error){
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		log.Printf("Error fetching task. %d: %v", taskId, err)
		return nil, errors.New("failed to retrieve task")
	}
	return &task, nil
}