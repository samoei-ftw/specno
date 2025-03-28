package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *gorm.DB, user User) (uint, error) {
	user.CreatedAt = time.Now()
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
