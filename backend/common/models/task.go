package models

import (
	"time"

	"github.com/samoei-ftw/specno/backend/common/enums"
)

type Task struct {
	ID          int       `json:"id"`
	UserID      uint      `json:"user_id"`
	ProjectID   uint      `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status 		enums.TaskStatus `json:"status"`
}
