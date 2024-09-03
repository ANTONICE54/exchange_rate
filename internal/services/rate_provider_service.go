package services

import (
	"rate/internal/apperrors"
	"rate/internal/pkg/provider"
)

//go:generate mockgen -source=rate_provider_service.go -destination=mock/rate_provider_service.go

type RateProviderService struct {
	rateProvider provider.IRateProvider
}

func NewRateProviderService(RateProvider provider.IRateProvider) *RateProviderService {
	return &RateProviderService{
		rateProvider: RateProvider,
	}
}

func (s *RateProviderService) Get() (*float64, error) {
	rate, err := s.rateProvider.GetRate()

	if err != nil {
		return nil, apperrors.NewError("Failed to fetch the rate", apperrors.Internal)
	}

	return rate, nil
}
