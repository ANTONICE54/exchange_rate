// Code generated by MockGen. DO NOT EDIT.
// Source: subscription_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "rate/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISubscriptionRepo is a mock of ISubscriptionRepo interface.
type MockISubscriptionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockISubscriptionRepoMockRecorder
}

// MockISubscriptionRepoMockRecorder is the mock recorder for MockISubscriptionRepo.
type MockISubscriptionRepoMockRecorder struct {
	mock *MockISubscriptionRepo
}

// NewMockISubscriptionRepo creates a new mock instance.
func NewMockISubscriptionRepo(ctrl *gomock.Controller) *MockISubscriptionRepo {
	mock := &MockISubscriptionRepo{ctrl: ctrl}
	mock.recorder = &MockISubscriptionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubscriptionRepo) EXPECT() *MockISubscriptionRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISubscriptionRepo) Create(email models.Email) (*models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", email)
	ret0, _ := ret[0].(*models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockISubscriptionRepoMockRecorder) Create(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISubscriptionRepo)(nil).Create), email)
}

// GetByID mocks base method.
func (m *MockISubscriptionRepo) GetByID(emailID uint) (*models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", emailID)
	ret0, _ := ret[0].(*models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockISubscriptionRepoMockRecorder) GetByID(emailID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockISubscriptionRepo)(nil).GetByID), emailID)
}

// List mocks base method.
func (m *MockISubscriptionRepo) List() ([]*models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockISubscriptionRepoMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockISubscriptionRepo)(nil).List))
}

// MockISubscriptionService is a mock of ISubscriptionService interface.
type MockISubscriptionService struct {
	ctrl     *gomock.Controller
	recorder *MockISubscriptionServiceMockRecorder
}

// MockISubscriptionServiceMockRecorder is the mock recorder for MockISubscriptionService.
type MockISubscriptionServiceMockRecorder struct {
	mock *MockISubscriptionService
}

// NewMockISubscriptionService creates a new mock instance.
func NewMockISubscriptionService(ctrl *gomock.Controller) *MockISubscriptionService {
	mock := &MockISubscriptionService{ctrl: ctrl}
	mock.recorder = &MockISubscriptionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubscriptionService) EXPECT() *MockISubscriptionServiceMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockISubscriptionService) Subscribe(email models.Email) (*models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", email)
	ret0, _ := ret[0].(*models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockISubscriptionServiceMockRecorder) Subscribe(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockISubscriptionService)(nil).Subscribe), email)
}
