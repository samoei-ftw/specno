package repo

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/samoei-ftw/specno/backend/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
type UserRepository interface {
	Create(user *models.User) error
	GetByUserID(userID int) ([]models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return nil
}
func CreateUser(user models.User) (uint, error) {
	// Initialize the database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var db_err error
	DB, db_err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db_err != nil {
		log.Fatal("Failed to connect to the database:", db_err)
	}
	db_err = DB.AutoMigrate(&models.User{})
	if db_err != nil {
		log.Fatal("Failed to migrate database:", db_err)
	}
	if DB == nil {
		return 0, errors.New("database connection is not initialized")
	}

	result := DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	var user models.User
	result := DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
