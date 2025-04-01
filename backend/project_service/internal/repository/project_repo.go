package repository

import (
	commonModels "common/internal/models"
	"errors"
	"fmt"
	"log"
	"os"

	taskoModels "tasko/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Interface
type ProjectRepository interface {
	Create(project *commonModels.Project) error
	GetByUserID(userID int) ([]commonModels.Project, error)
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

	err = DB.AutoMigrate(&taskoModels.User{}, &commonModels.Project{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return err
	}

	return nil
}

func (r *projectRepo) Create(project *commonModels.Project) error {
	if DB == nil {
		return errors.New("DB connection error.")
	}
	return r.db.Create(project).Error
}

func (r *projectRepo) GetByUserID(userID int) ([]commonModels.Project, error) {
	if DB == nil {
		return nil, errors.New("DB connection error.")
	}
	var projects []commonModels.Project
	if err := r.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
