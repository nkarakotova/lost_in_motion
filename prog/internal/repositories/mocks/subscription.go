// Code generated by MockGen. DO NOT EDIT.
// Source: subscription.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	models "prog/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSubscriptionRepository is a mock of SubscriptionRepository interface.
type MockSubscriptionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionRepositoryMockRecorder
}

// MockSubscriptionRepositoryMockRecorder is the mock recorder for MockSubscriptionRepository.
type MockSubscriptionRepositoryMockRecorder struct {
	mock *MockSubscriptionRepository
}

// NewMockSubscriptionRepository creates a new mock instance.
func NewMockSubscriptionRepository(ctrl *gomock.Controller) *MockSubscriptionRepository {
	mock := &MockSubscriptionRepository{ctrl: ctrl}
	mock.recorder = &MockSubscriptionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionRepository) EXPECT() *MockSubscriptionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSubscriptionRepository) Create(ctx context.Context, subscription *models.Subscription, clientID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, subscription, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSubscriptionRepositoryMockRecorder) Create(ctx, subscription, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubscriptionRepository)(nil).Create), ctx, subscription, clientID)
}

// GetByID mocks base method.
func (m *MockSubscriptionRepository) GetByID(ctx context.Context, id uint64) (*models.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockSubscriptionRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSubscriptionRepository)(nil).GetByID), ctx, id)
}

// IncreaseRemainingTrainingsNum mocks base method.
func (m *MockSubscriptionRepository) IncreaseRemainingTrainingsNum(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseRemainingTrainingsNum", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseRemainingTrainingsNum indicates an expected call of IncreaseRemainingTrainingsNum.
func (mr *MockSubscriptionRepositoryMockRecorder) IncreaseRemainingTrainingsNum(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseRemainingTrainingsNum", reflect.TypeOf((*MockSubscriptionRepository)(nil).IncreaseRemainingTrainingsNum), ctx, id)
}

// ReduceRemainingTrainingsNum mocks base method.
func (m *MockSubscriptionRepository) ReduceRemainingTrainingsNum(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceRemainingTrainingsNum", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReduceRemainingTrainingsNum indicates an expected call of ReduceRemainingTrainingsNum.
func (mr *MockSubscriptionRepositoryMockRecorder) ReduceRemainingTrainingsNum(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceRemainingTrainingsNum", reflect.TypeOf((*MockSubscriptionRepository)(nil).ReduceRemainingTrainingsNum), ctx, id)
}
