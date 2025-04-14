package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	UserID      uint      `json:"user_id"`
	ProjectID   uint      `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status 		string 	  `json:"status"`
}
