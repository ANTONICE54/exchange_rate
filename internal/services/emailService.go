package services

import (
	"rate/internal/database"
	"rate/internal/pkg/mailer"
	"rate/internal/pkg/provider"
	"sync"
)

type EmailService struct {
	mailServer   mailer.IMailer
	emailRepo    database.IEmailRepo
	rateProvider provider.IRateProvider
	wg           *sync.WaitGroup
}

func NewEmailService(mailServer mailer.IMailer, emailRepo database.IEmailRepo, rateProvider provider.IRateProvider, wg *sync.WaitGroup) *EmailService {
	return &EmailService{
		mailServer:   mailServer,
		emailRepo:    emailRepo,
		rateProvider: rateProvider,
		wg:           wg,
	}
}

func (s *EmailService) SendEmails() {
	emails, err := s.emailRepo.ListEmails()
	if err != nil {
		return
	}

	rate, err := s.rateProvider.GetRate()
	if err != nil {
		return
	}

	for _, e := range emails {
		s.wg.Add(1)
		go s.sendEmail(e.Email, *rate)
	}
}

func (s *EmailService) sendEmail(email string, rate float64) {
	defer s.wg.Done()
	s.mailServer.SendEmail(email, rate)
}
