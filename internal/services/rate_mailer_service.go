package services

import (
	"rate/internal/pkg/provider"
	"sync"
)

//go:generate mockgen -source=rate_mailer_service.go -destination=mock/rate_mailer_service.go

type IMailer interface {
	SendEmail(email string, rate float64)
}

type RateMailerService struct {
	mailServer   IMailer
	emailRepo    ISubscriptionRepo
	rateProvider provider.IRateProvider
	wg           *sync.WaitGroup
}

func NewRateMailerService(mailServer IMailer, emailRepo ISubscriptionRepo, rateProvider provider.IRateProvider, wg *sync.WaitGroup) *RateMailerService {
	return &RateMailerService{
		mailServer:   mailServer,
		emailRepo:    emailRepo,
		rateProvider: rateProvider,
		wg:           wg,
	}
}

func (s *RateMailerService) SendEmails() {
	emails, err := s.emailRepo.List()
	if err != nil {
		return
	}

	rate, err := s.rateProvider.GetRate()
	if err != nil {
		return
	}
	semaphore := make(chan struct{}, 5)
	for _, e := range emails {

		s.wg.Add(1)
		semaphore <- struct{}{}
		go func(email string) {
			defer s.wg.Done()
			defer func() { <-semaphore }()

			s.mailServer.SendEmail(email, *rate)

		}(e.Email)
	}
	s.wg.Wait()
}
