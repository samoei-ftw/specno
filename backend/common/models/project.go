package models

import (
	"time"
)
type Project struct {
	ID          uint   `json:"id"`
	UserID 		uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"descriptions"`
	CreatedAt   time.Time `json:"created_at"`
}
