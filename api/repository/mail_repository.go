package repository

import (
	"errors"
	"log"
	"log/slog"
	"mime"
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

type IMailRepository interface {
	SendMail(to string, subject string, body string) error
}

type MailRepository struct {
}

func NewMailRepository() IMailRepository {
	return &MailRepository{}
}

func (m *MailRepository) SendMail(to string, subject string, body string) error {
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
