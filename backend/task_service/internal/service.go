package internal

import (
	"github.com/samoei-ftw/specno/backend/common/models"
	//"github.com/samoei-ftw/specno/backend/gateways"
	"github.com/samoei-ftw/specno/backend/common/enums"
)

var (
	userGateway = gateways.UserGatewayInit()
	//projectGateway = gateways.ProjectGatewayInit() TODO
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
	project, err := projectGateway.userOwnsProject(userId, projectId)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	project := &models.Project{
		Name:        name,
		Description: description,
		UserID:      user.ID,
	}

	if err := s.repo.Create(project); err != nil {
		log.Printf("Error creating project: %v", err)
		return nil, errors.New("failed to create project")
	}*/

	return project, nil
}

// Lists all tasks for a project
func (s *service) ListTasks(projectId uint) ([]*models.Task, error) {
	/** Validate user exists
	user, err := userGateway.GetUserByID(userId)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userId, err)
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	projectList, err := s.repo.GetByUserID(user.ID)
	if err != nil {
		log.Printf("Error fetching projects for user %d: %v", userId, err)
		return nil, errors.New("failed to retrieve projects")
	}

	projects := make([]*models.Project, len(projectList))
	for i, p := range projectList {
		projects[i] = &p
	}

	return projects, nil*/
}

func (s *service) GetProject(projectId uint) (*models.Project, error){
	/**project, err := s.repo.GetProjectById(projectId)
	if err != nil {
		log.Printf("Error fetching project. %d: %v", projectId, err)
		return nil, errors.New("failed to retrieve project")
	}
	return &project, nil*/
}