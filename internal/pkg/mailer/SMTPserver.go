package mailer

import (
	"fmt"
	"log"
	"net/smtp"
)

type IMailer interface {
	SendEmail(email string, rate float64)
}

type SMTPServer struct {
	host     string
	port     string
	username string
	password string
	from     string
	auth     smtp.Auth
}

func NewSMTPServer(host, port, username, password, from string) *SMTPServer {

	server := &SMTPServer{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
	auth := smtp.PlainAuth("", server.username, server.password, server.host)

	server.auth = auth

	return server
}

func (s *SMTPServer) SendEmail(email string, rate float64) {
	subject := "Subject: Daily Exchange Rate\n"
	fromHeader := fmt.Sprintf("From: %s\n", s.from)
	body := fmt.Sprintf("The current exchange rate is: %f UAH/USD", rate)

	toHeader := fmt.Sprintf("To: %s\n", email)

	msg := []byte(subject + fromHeader + toHeader + "\n" + body)
	err := smtp.SendMail(s.host+":"+s.port, s.auth, s.from, []string{email}, msg)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Email sent successfuly!")

}
