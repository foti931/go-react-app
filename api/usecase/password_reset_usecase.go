package usecase

import (
	"errors"
	"go-rest-api/models"
	"go-rest-api/repository"
)

type IPasswordResetUseCase interface {
	GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error)
	ResetPassword(request *models.PasswordReset) (*models.User, error)
}

type passwordResetUseCase struct {
	pr repository.IPasswordRepository
}

func (p *passwordResetUseCase) GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error) {
	return p.pr.GetPasswordResetRequest(request)
}

func (p *passwordResetUseCase) ResetPassword(request *models.PasswordReset) (*models.User, error) {
	return nil, errors.New("not implemented")
}

func NewPasswordResetUseCase(pr repository.IPasswordRepository) IPasswordResetUseCase {
	return &passwordResetUseCase{pr: pr}
}
