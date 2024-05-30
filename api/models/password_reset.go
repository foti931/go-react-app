package models

import "gorm.io/gorm"

type PasswordReset struct {
	*gorm.Model
	Email  string `json:"email"`
	Token  string `json:"token"`
	User   User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId uint   `json:"user_id" gorm:"not null"`
}

type PasswordResetRequestResponse struct {
	Email string `json:"email"`
}
