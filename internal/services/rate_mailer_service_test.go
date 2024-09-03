package services

import (
	"rate/internal/models"
	mock "rate/internal/services/mock"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestRateMailerService(t *testing.T) {
	ctrl := gomock.NewController(t)

	subscRepoMock := mock.NewMockISubscriptionRepo(ctrl)
	mailerMock := mock.NewMockIMailer(ctrl)
	rateProviderMock := mock.NewMockIRateProvider(ctrl)
	emailList := []*models.Email{
		{ID: 1,
			Email:     "test1@example.com",
			CreatedAt: time.Now(),
		},
		{ID: 2,
			Email:     "test2@example.com",
			CreatedAt: time.Now(),
		},
		{ID: 3,
			Email:     "test3@example.com",
			CreatedAt: time.Now(),
		},
	}

	rate := 54.52
	subscRepoMock.EXPECT().List().Return(emailList, nil).Times(1)
	rateProviderMock.EXPECT().GetRate().Return(&rate, nil).Times(1)
	for _, e := range emailList {
		mailerMock.EXPECT().SendEmail(e.Email, rate).Return().Times(1)
	}

	mailerService := NewRateMailerService(mailerMock, subscRepoMock, rateProviderMock, &sync.WaitGroup{})

	mailerService.SendEmails()

}
