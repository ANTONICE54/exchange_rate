package services

import (
	"rate/internal/apperrors"
	"rate/internal/models"
	mock "rate/internal/services/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionService(t *testing.T) {

	ctrl := gomock.NewController(t)

	mockSubscRepo := mock.NewMockISubscriptionRepo(ctrl)
	subscriptionService := NewSubscriptionService(mockSubscRepo)
	createdAt := time.Now()

	testCases := []struct {
		name          string
		mockFunc      func()
		passedEmail   string
		expectedEmail *models.Email
		expectedError error
	}{
		{
			name: "Successful subscription",
			mockFunc: func() {

				email := models.Email{
					ID:        1,
					Email:     "test@gmail.com",
					CreatedAt: createdAt,
				}

				mockSubscRepo.EXPECT().Create(models.Email{Email: "test@gmail.com"}).Return(&email, nil).Times(1)
			},
			passedEmail: "test@gmail.com",
			expectedEmail: &models.Email{
				ID:        1,
				Email:     "test@gmail.com",
				CreatedAt: createdAt,
			},
			expectedError: nil,
		},
		{
			name: "Email is already subscribed",
			mockFunc: func() {
				err := &pq.Error{
					Code: "23505",
				}

				mockSubscRepo.EXPECT().Create(models.Email{Email: "test@gmail.com"}).Return(nil, err).Times(1)
			},
			passedEmail:   "test@gmail.com",
			expectedEmail: nil,
			expectedError: apperrors.NewError("Subscription with such email already exists", apperrors.BadRequest),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			receivedeEmail, receivedError := subscriptionService.Subscribe(models.Email{Email: tc.passedEmail})

			assert.Equal(t, receivedError, tc.expectedError)
			if tc.expectedEmail != nil {
				assert.Equal(t, *receivedeEmail, *tc.expectedEmail)
			} else {
				assert.Nil(t, receivedeEmail)
			}
		})
	}

}
