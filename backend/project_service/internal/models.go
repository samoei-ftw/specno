package internal

import (
	"gorm.io/gorm"
)
type projectRepo struct {
	db *gorm.DB
}
type ProjectService struct {
	repo Repository
}
