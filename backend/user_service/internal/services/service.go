package services

import (
	"errors"
	"log"

	"github.com/samoei-ftw/specno/backend/common/models"
	dto "github.com/samoei-ftw/specno/backend/user_service/internal/models"
	"github.com/samoei-ftw/specno/backend/user_service/internal/repo"

	"golang.org/x/crypto/bcrypt"
)
type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}
// Fetches a user by their email.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
    user, err := s.repo.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// RegisterUser handles user registration.
func (s *UserService) RegisterUser(email, password string) (uint, error) {
	// Check if user already exists
	existingUser, _ := s.repo.GetUserByEmail(email)
	if existingUser != nil {
		return 0, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	userID, err := s.repo.Create(&user)
	if err != nil {
		log.Println("Error saving user to DB:", err)
		return 0, err
	}

	return userID, nil
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpsertUser(upsertUser dto.UpsertUser) (*models.User, error) {
	user := &models.User{}
	if upsertUser.UserID == nil {
		return nil, errors.New("user ID is required for update")
	}
	if upsertUser.UserID != nil {
		existingUser, err := s.repo.GetUserByID(int(*upsertUser.UserID))
		if err != nil {
			return nil, err
		}

		if existingUser != nil {
			user = existingUser
		}
	}
	// TODO: validation
	if upsertUser.Email != nil {
		user.Email = *upsertUser.Email
	}
	if upsertUser.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*upsertUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	}

	return s.repo.Upsert(user)

}