package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rate/internal/apperrors"
	"rate/internal/models"
	mock "rate/internal/server/handlers/mock"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	mockSubscService := mock.NewMockISubscriptionService(ctrl)
	subscHandler := NewSubscriptionHandler(mockSubscService)

	router.POST("/subscribe", subscHandler.Subscribe)

	email := "test1@test.com"
	createdAt := time.Now()

	testCases := []struct {
		name                 string
		mockFunc             func()
		requestBodyFunc      func() ([]byte, error)
		expectedCode         int
		expectedResponseBody func() ([]byte, error)
	}{
		{
			name: "Success",
			mockFunc: func() {
				mockSubscService.EXPECT().Subscribe(models.Email{Email: email}).Return(&models.Email{ID: 1, Email: email, CreatedAt: createdAt}, nil).Times(1)
			},
			requestBodyFunc: func() ([]byte, error) {
				subscData := subscribeInput{
					Email: email,
				}
				return json.Marshal(subscData)
			},

			expectedCode: http.StatusOK,
			expectedResponseBody: func() ([]byte, error) {
				return nil, nil
			},
		},
		{
			name: "Bad request",
			mockFunc: func() {
				mockSubscService.EXPECT().Subscribe(models.Email{Email: email}).Return(&models.Email{ID: 1, Email: email, CreatedAt: createdAt}, nil).Times(0)
			},
			requestBodyFunc: func() ([]byte, error) {
				return []byte("bad request"), nil
			},

			expectedCode: http.StatusBadRequest,
			expectedResponseBody: func() ([]byte, error) {
				return json.Marshal(map[string]string{"error": "Bad request"})
			},
		},
		{
			name: "Internal error",
			mockFunc: func() {
				mockSubscService.EXPECT().Subscribe(models.Email{Email: email}).Return(nil, apperrors.ErrDatabase).Times(1)
			},
			requestBodyFunc: func() ([]byte, error) {
				subscData := subscribeInput{
					Email: email,
				}
				return json.Marshal(subscData)
			},

			expectedCode: http.StatusInternalServerError,
			expectedResponseBody: func() ([]byte, error) {
				return json.Marshal(map[string]string{"error": "Database raised an error"})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()

			requestBody, err := tc.requestBodyFunc()
			assert.NoError(t, err)

			respRecorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/subscribe", bytes.NewBuffer(requestBody))
			assert.NoError(t, err)

			router.ServeHTTP(respRecorder, request)

			expectedResponseBody, err := tc.expectedResponseBody()
			assert.NoError(t, err)

			assert.Equal(t, expectedResponseBody, respRecorder.Body.Bytes())
			assert.Equal(t, tc.expectedCode, respRecorder.Code)

		})
	}

}
