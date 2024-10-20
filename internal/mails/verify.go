package mails

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"strconv"
)

func (m *Mail) SendEmail(to string, message string) error {
	var (
		Host     = os.Getenv("SMTP_HOST")
		Port, _  = strconv.Atoi(os.Getenv("SMTP_PORT"))
		Username = os.Getenv("SMTP_USERNAME")
		Pass     = os.Getenv("SMTP_PASSWORD")
	)
	mail := gomail.NewMessage()
	mail.SetHeader("From", "Support Team <email@example.com>")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "Verify your email")
	mail.SetBody("text/html", message)
	d := gomail.NewDialer(Host, Port, Username, Pass)
	return d.DialAndSend(mail)
}
func (m *Mail) VerifyEmail(pass string) bool {
	return m.Code == pass
}
func (*Mail) GenerateCode() string {
	return strconv.Itoa(int(rand.Int31()))
}
