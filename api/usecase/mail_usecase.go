package usecase

import "go-rest-api/repository"

type IMailUsecase interface {
	SendMail(to string, subject string, token string, body string) error
}

type MailUsecase struct {
	mr repository.IMailRepository
}

func NewMailUsecase(mr repository.IMailRepository) IMailUsecase {
	return &MailUsecase{mr: mr}
}

// SendMail implements IMailInterface.
func (mu *MailUsecase) SendMail(to string, subject string, token string, body string) error {
	return mu.mr.SendMail(to, subject, body)
}
