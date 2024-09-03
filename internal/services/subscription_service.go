package services

import (
	"rate/internal/apperrors"
	"rate/internal/models"

	"github.com/lib/pq"
)

type ISubscriptionRepo interface {
	Create(email models.Email) (*models.Email, error)
	List() ([]*models.Email, error)
	GetByID(emailID uint) (*models.Email, error)
}

//go:generate mockgen -source=subscription_service.go -destination=mock/subscription_service.go

type ISubscriptionService interface {
	Subscribe(email models.Email) (*models.Email, error)
}

type SubscriptionService struct {
	subscRepo ISubscriptionRepo
}

func NewSubscriptionService(subscRepo ISubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{
		subscRepo: subscRepo,
	}
}

func (s SubscriptionService) Subscribe(email models.Email) (*models.Email, error) {
	subsc, err := s.subscRepo.Create(email)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return nil, apperrors.NewError("Subscription with such email already exists", apperrors.BadRequest)
		} else {
			return nil, apperrors.ErrDatabase
		}
	}

	return subsc, nil
}
