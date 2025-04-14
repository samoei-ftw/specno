package internal

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type UpdateTaskRequest struct {
	Status int `json:"status" validate:"required"`
}