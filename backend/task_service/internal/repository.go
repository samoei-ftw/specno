package internal

import (
	"errors"

	"github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(task *models.Task) error {
	if r.db == nil {
		return errors.New("DB connection not initialized")
	}
	return r.db.Create(task).Error
}