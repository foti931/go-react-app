package repository

import (
	"go-rest-api/models"
	"log/slog"

	"gorm.io/gorm"
)

type IPasswordRepository interface {
	CreatePasswordResetRequest(user *models.User, token string) error
	GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error)
	DeleteAllPasswordResetRequest(request *models.PasswordReset) error
}

type PasswordRepository struct {
	db *gorm.DB
}

func (p *PasswordRepository) DeleteAllPasswordResetRequest(request *models.PasswordReset) error {
	if err := p.db.Where("user_id = ?", request.UserId).Delete(&models.PasswordReset{}); err != nil {
		return err.Error
	}
	return nil
}

func (p *PasswordRepository) GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error) {
	user := &models.User{}

	if err := p.db.Joins("INNER JOIN users ON users.id = password_resets.user_id").Where("password_resets.email = ? AND password_resets.token = ?", request.Email, request.Token).Order("password_resets.id DESC").Last(&request).Error; err != nil {
		return nil, err
	}

	if err := p.db.Where("email = ?", request.Email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// CreatePasswordResetRequest パスワード変更リクエストを作成する
func (p *PasswordRepository) CreatePasswordResetRequest(user *models.User, token string) error {

	request := &models.PasswordReset{
		Email:  user.Email,
		Token:  token,
		UserId: user.ID,
	}

	result := p.db.Create(request)
	if result.RowsAffected == 0 {
		slog.Info(result.Error.Error())
		return result.Error
	}

	return nil
}

func NewPasswordRepository(db *gorm.DB) IPasswordRepository {
	return &PasswordRepository{db: db}
}
