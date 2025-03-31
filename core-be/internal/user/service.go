package user

import (
	"errors"

	"tasko/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// Creates a new user
func RegisterUser(email, password string) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return CreateUser(user)
}
