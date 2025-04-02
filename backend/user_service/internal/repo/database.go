package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/samoei-ftw/specno/backend/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDatabase sets up the database connection and runs migrations
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

	// Run migrations
	if err = DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}

	log.Println("✅ Database connection successfully established")
	return nil
}

// GetDB returns the database instance for repo functions
func GetDB() *gorm.DB {
	return DB
}