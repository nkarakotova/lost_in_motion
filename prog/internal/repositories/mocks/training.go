// Code generated by MockGen. DO NOT EDIT.
// Source: training.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	models "prog/internal/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockTrainingRepository is a mock of TrainingRepository interface.
type MockTrainingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrainingRepositoryMockRecorder
}

// MockTrainingRepositoryMockRecorder is the mock recorder for MockTrainingRepository.
type MockTrainingRepositoryMockRecorder struct {
	mock *MockTrainingRepository
}

// NewMockTrainingRepository creates a new mock instance.
func NewMockTrainingRepository(ctrl *gomock.Controller) *MockTrainingRepository {
	mock := &MockTrainingRepository{ctrl: ctrl}
	mock.recorder = &MockTrainingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrainingRepository) EXPECT() *MockTrainingRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTrainingRepository) Create(ctx context.Context, training *models.Training) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, training)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTrainingRepositoryMockRecorder) Create(ctx, training interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrainingRepository)(nil).Create), ctx, training)
}

// Delete mocks base method.
func (m *MockTrainingRepository) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTrainingRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrainingRepository)(nil).Delete), ctx, id)
}

// GetAllBetweenDateTime mocks base method.
func (m *MockTrainingRepository) GetAllBetweenDateTime(ctx context.Context, start, end time.Time) ([]models.Training, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBetweenDateTime", ctx, start, end)
	ret0, _ := ret[0].([]models.Training)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBetweenDateTime indicates an expected call of GetAllBetweenDateTime.
func (mr *MockTrainingRepositoryMockRecorder) GetAllBetweenDateTime(ctx, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBetweenDateTime", reflect.TypeOf((*MockTrainingRepository)(nil).GetAllBetweenDateTime), ctx, start, end)
}

// GetAllByClient mocks base method.
func (m *MockTrainingRepository) GetAllByClient(ctx context.Context, id uint64) ([]models.Training, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByClient", ctx, id)
	ret0, _ := ret[0].([]models.Training)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByClient indicates an expected call of GetAllByClient.
func (mr *MockTrainingRepositoryMockRecorder) GetAllByClient(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByClient", reflect.TypeOf((*MockTrainingRepository)(nil).GetAllByClient), ctx, id)
}

// GetAllByCoachOnDate mocks base method.
func (m *MockTrainingRepository) GetAllByCoachOnDate(ctx context.Context, id uint64, date time.Time) ([]models.Training, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByCoachOnDate", ctx, id, date)
	ret0, _ := ret[0].([]models.Training)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByCoachOnDate indicates an expected call of GetAllByCoachOnDate.
func (mr *MockTrainingRepositoryMockRecorder) GetAllByCoachOnDate(ctx, id, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByCoachOnDate", reflect.TypeOf((*MockTrainingRepository)(nil).GetAllByCoachOnDate), ctx, id, date)
}

// GetAllByDateTime mocks base method.
func (m *MockTrainingRepository) GetAllByDateTime(ctx context.Context, dateTime time.Time) ([]models.Training, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByDateTime", ctx, dateTime)
	ret0, _ := ret[0].([]models.Training)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByDateTime indicates an expected call of GetAllByDateTime.
func (mr *MockTrainingRepositoryMockRecorder) GetAllByDateTime(ctx, dateTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByDateTime", reflect.TypeOf((*MockTrainingRepository)(nil).GetAllByDateTime), ctx, dateTime)
}

// GetByID mocks base method.
func (m *MockTrainingRepository) GetByID(ctx context.Context, id uint64) (*models.Training, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Training)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTrainingRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTrainingRepository)(nil).GetByID), ctx, id)
}

// IncreaseAvailablePlacesNum mocks base method.
func (m *MockTrainingRepository) IncreaseAvailablePlacesNum(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseAvailablePlacesNum", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseAvailablePlacesNum indicates an expected call of IncreaseAvailablePlacesNum.
func (mr *MockTrainingRepositoryMockRecorder) IncreaseAvailablePlacesNum(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseAvailablePlacesNum", reflect.TypeOf((*MockTrainingRepository)(nil).IncreaseAvailablePlacesNum), ctx, id)
}

// ReduceAvailablePlacesNum mocks base method.
func (m *MockTrainingRepository) ReduceAvailablePlacesNum(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceAvailablePlacesNum", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReduceAvailablePlacesNum indicates an expected call of ReduceAvailablePlacesNum.
func (mr *MockTrainingRepositoryMockRecorder) ReduceAvailablePlacesNum(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceAvailablePlacesNum", reflect.TypeOf((*MockTrainingRepository)(nil).ReduceAvailablePlacesNum), ctx, id)
}
