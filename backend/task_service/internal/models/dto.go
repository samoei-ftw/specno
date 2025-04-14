package internal

import "github.com/samoei-ftw/specno/backend/common/enums"

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type UpdateTaskRequest struct {
	Status enums.TaskStatus `json:"status" validate:"required"`
}