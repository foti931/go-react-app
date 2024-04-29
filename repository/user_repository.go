package repository

import (
	"go-rest-api/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) GetUserByEmail(user *models.User, email string) error {
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) CreateUser(user *models.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
