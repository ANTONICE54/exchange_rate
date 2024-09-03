package services

import (
	"rate/internal/apperrors"
	mock "rate/internal/services/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRateProvider(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRateProvider := mock.NewMockIRateProvider(ctrl)
	rateService := NewRateProviderService(mockRateProvider)

	rate := 52.54
	testCases := []struct {
		name          string
		mockFunc      func()
		expectedRate  *float64
		expectedError error
	}{{
		name: "Rate fetched successfully",
		mockFunc: func() {
			rateToReturn := 52.54
			mockRateProvider.EXPECT().GetRate().Return(&rateToReturn, nil).Times(1)
		},
		expectedRate:  &rate,
		expectedError: nil,
	},
		{
			name: "Failed to fetch rate",
			mockFunc: func() {
				errToReturn := apperrors.NewError("Failed to fetch the rate", apperrors.Internal)
				mockRateProvider.EXPECT().GetRate().Return(nil, errToReturn).Times(1)
			},
			expectedRate:  nil,
			expectedError: apperrors.NewError("Failed to fetch the rate", apperrors.Internal),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			recievedRate, err := rateService.Get()
			require.Equal(t, err, tc.expectedError)
			if tc.expectedRate == nil {
				require.Nil(t, recievedRate)
			} else {
				require.Equal(t, *tc.expectedRate, *recievedRate)
			}
		})
	}

}
