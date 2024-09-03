package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rate/internal/apperrors"
	mock "rate/internal/server/handlers/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRateHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockRateService := mock.NewMockIRateProviderService(ctrl)
	rateHandler := NewRateHandler(mockRateService)

	router := gin.Default()
	router.GET("/rate", rateHandler.Get)
	request, err := http.NewRequest(http.MethodGet, "/rate", nil)
	assert.NoError(t, err)

	testCases := []struct {
		name                 string
		mockFunc             func()
		expectedResponseBody func() ([]byte, error)
		expectedCode         int
	}{{
		name: "Success",
		mockFunc: func() {
			rate := 52.54
			mockRateService.EXPECT().Get().Return(&rate, nil).Times(1)
		},
		expectedResponseBody: func() ([]byte, error) {
			return []byte("52.54"), nil
		},
		expectedCode: http.StatusOK,
	},
		{
			name: "Internal server error",
			mockFunc: func() {

				mockRateService.EXPECT().Get().Return(nil, apperrors.NewError("Failed to fetch the rate", apperrors.Internal)).Times(1)
			},
			expectedResponseBody: func() ([]byte, error) {
				return json.Marshal(map[string]string{"error": "Failed to fetch the rate"})
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			respRecorer := httptest.NewRecorder()
			router.ServeHTTP(respRecorer, request)
			expectedResponseBody, err := tc.expectedResponseBody()
			assert.NoError(t, err)
			assert.Equal(t, respRecorer.Body.Bytes(), expectedResponseBody)
			assert.Equal(t, respRecorer.Code, tc.expectedCode)

		})
	}

}
