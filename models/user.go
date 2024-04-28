package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
