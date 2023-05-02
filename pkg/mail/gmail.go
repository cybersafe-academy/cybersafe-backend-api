package mail

import (
	"cybersafe-backend-api/pkg/settings"
	"net/smtp"
)

const (
	DefaultForgotPasswordSubject = "CyberSafe Academy - Reset your password"
)

type GmailMailer struct {
	Email    string
	Password string
}

func Config(config settings.Settings) Mailer {
	return &GmailMailer{
		Email:    config.String("mail.email"),
		Password: config.String("mail.password"),
	}
}

func (gm *GmailMailer) Send(to []string, subject string, message string) error {
	auth := smtp.PlainAuth("", gm.Email, gm.Password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, gm.Email, to, []byte(message))

	if err != nil {
		return err
	}

	return nil
}
