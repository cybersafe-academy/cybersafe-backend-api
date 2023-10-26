package mail

import (
	"cybersafe-backend-api/pkg/settings"
	"log"
	"net/smtp"
	"strings"
)

const (
	DefaultForgotPasswordSubject = "CyberSafe Academy - Reset your password"
	DefaultFirstAccessSubject    = "CyberSafe Academy - Finish your signup"
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
	headers := make(map[string]string)
	headers["From"] = gm.Email
	headers["To"] = strings.Join(to, ",")
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	body := ""

	for key, value := range headers {
		body += key + ": " + value + "\r\n"
	}

	body += "\r\n" + message
	auth := smtp.PlainAuth("", gm.Email, gm.Password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, gm.Email, to, []byte(body))

	if err != nil {
		log.Println("Error sending email:", err)

		return err
	}

	return nil
}
