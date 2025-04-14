package internal

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type UpdateTaskRequest struct {
	UserID *uint `json:"user_id" validate:"required"`
	Status *int `json:"status" validate:"required"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}