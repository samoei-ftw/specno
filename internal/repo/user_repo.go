package repo

import (
	"tasko/internal/models"

	"gorm.io/gorm"
)

// Get a user by their email
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
