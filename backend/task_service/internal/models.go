package internal

import (
	"gorm.io/gorm"
)
var taskCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint    `json:"user_id"`
}
type repo struct {
	db *gorm.DB
}