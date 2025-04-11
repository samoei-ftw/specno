package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
// sets up the database connection
func InitializeDatabase() error {
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
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	return nil
}


func RunMigrations(migrationsDir string) error {
	log.Println("Running migrations...")
	var err error
	if err = DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}

	log.Println("Database connection successfully established.")
	return nil
}

// returns the database instance for repo functions
func GetDB() *gorm.DB {
	return DB
}