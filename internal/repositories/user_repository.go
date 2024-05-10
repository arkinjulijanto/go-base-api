package repositories

import (
	"github.com/arkinjulijanto/go-base-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(u *URConfig) UserRepository {
	return &userRepository{
		db: u.DB,
	}
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByUsername(username string) (*models.User, error) {
	var user *models.User

	result := u.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
