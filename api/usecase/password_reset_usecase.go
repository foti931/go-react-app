package usecase

import (
	"go-rest-api/models"
	"go-rest-api/repository"
)

type IPasswordResetUseCase interface {
	GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error)
}

type passwordResetUseCase struct {
	pr repository.IPasswordRespository
}

func (p *passwordResetUseCase) GetPasswordResetRequest(request *models.PasswordReset) (*models.User, error) {
	return p.pr.GetPasswordResetRequest(request)
}

func NewPasswordResetUseCase(pr repository.IPasswordRespository) IPasswordResetUseCase {
	return &passwordResetUseCase{pr: pr}
}
