package internal

import (
	"github.com/samoei-ftw/specno/backend/common/models"
)
type Repository interface {
	Create(task *models.Task) error
}