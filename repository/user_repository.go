package repository

import "go-rest-api/models"

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetUserByEmail(user *models.User, email string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) CreateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}
