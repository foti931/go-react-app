package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model `json:"*_gorm_._model,omitempty"`
	Title       string `json:"title" gorm:"not null"`
	User        User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint   `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
