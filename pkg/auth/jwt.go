package services

import (
	"errors"
	"tasko/internal/models"
	"tasko/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(email, password string) (uint, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// save user
	userID, err := repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
