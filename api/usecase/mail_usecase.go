package usecase

import (
	"errors"
	"github.com/wneessen/go-mail"
	"log"
	"log/slog"
	"mime"
	"os"
	"strconv"
)

type IMailUsecase interface {
	SendMail(to string, subject string, token string, body string) error
}

type MailUsecase struct {
}

func NewMailUsecase() IMailUsecase {
	return &MailUsecase{}
}

// SendMail implements IMailInterface.
func (mu *MailUsecase) SendMail(to string, subject string, token string, body string) error {
	msg := mail.NewMsg()
	host := os.Getenv("SMTP_HOST")

	if host == "" {
		slog.Info("cannot get mail host")
		return errors.New("cannot get mail host")
	}

	var port int
	var err error
	if os.Getenv("SMTP_PORT") != "" {
		port, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
		if err != nil {
			log.Fatal(err)
		}
	}

	msg.From("admin@example.com")
	msg.AddTo(to)

	msg.Subject(mime.BEncoding.Encode("UTF-8", subject))
	msg.SetBodyString(mail.TypeTextPlain, body)

	c, err := mail.NewClient(host, mail.WithPort(port))
	c.SetTLSPolicy(mail.TLSOpportunistic)
	if err != nil {
		return errors.New("cannot create mail client")
	}

	if err := c.DialAndSend(msg); err != nil {
		slog.Error(err.Error())
		return errors.New("cannot send mail")
	}

	return nil
}
