package utils

import (
	"net/smtp"
	"usd-uah-testcase/internal"
)

type EmailProvider interface {
	SendMail(to string, subject string, body string) error
}

type SMTPProvider struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (s *SMTPProvider) SendMail(to string, subject string, body string) error {
	msg := "From: " + s.Username + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, s.Username, []string{to}, []byte(msg))
	return err
}

func NewEmailProvider() EmailProvider {
	config := internal.LoadConfig()
	return &SMTPProvider{
		Host:     config.STMPHost,
		Port:     config.STMPPort,
		Username: config.STMPUsername,
		Password: config.STMPPassword,
	}
}

var emailProvider EmailProvider = NewEmailProvider()

func SendMail(to string, subject string, body string) error {
	return emailProvider.SendMail(to, subject, body)
}
