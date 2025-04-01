package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	projectModels "project/internal/models"

	userModels "user/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Interface
type ProjectRepository interface {
	Create(project *projectModels.Project) error
	GetByUserID(userID int) ([]projectModels.Project, error)
}

// Struct
type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepo{db: db}
}

func InitializeDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return err
	}

	err = DB.AutoMigrate(&userModels.User{}, &projectModels.Project{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return err
	}

	return nil
}

func (r *projectRepo) Create(project *projectModels.Project) error {
	if DB == nil {
		return errors.New("DB connection error.")
	}
	return r.db.Create(project).Error
}

func (r *projectRepo) GetByUserID(userID int) ([]projectModels.Project, error) {
	if DB == nil {
		return nil, errors.New("DB connection error.")
	}
	var projects []projectModels.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
