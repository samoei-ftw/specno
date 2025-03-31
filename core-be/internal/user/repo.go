package user

import (
	"tasko/internal/models"
)

func CreateUser(user models.User) (uint, error) {
	result := models.DB.Create(&user) // Access the DB initialized in models
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
