package repo

import (
	"errors"

	"github.com/samoei-ftw/specno/backend/common/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

type UserRepository interface {
	Create(user *models.User) (uint, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) Create(user *models.User) (uint, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

func (u *userRepo) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	result := u.db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
