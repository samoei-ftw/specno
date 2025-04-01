package services

import (
	"errors"
	"log"
	"user/internal/models"
	"user/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

// Fetches a user by their email.
func GetUserByEmail(email string) (*models.User, error) {
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// RegisterUser handles user registration
func RegisterUser(email, password string) (uint, error) {
	// Check if user already exists
	//existingUser, _ := repo.GetUserByEmail(email)
	//if existingUser != nil {
	//	return 0, errors.New("user already exists")
	//}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	// Create a new user object
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// Call the repo layer to save the user
	userID, err := repo.CreateUser(user)
	if err != nil {
		log.Println("Error saving user to DB:", err)
		return 0, err
	}

	return userID, nil
}
