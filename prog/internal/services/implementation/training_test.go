package servicesImplementation

import (
	"os"
	"prog/internal/models"
	"prog/pkg/errors/servicesErrors"
	"prog/internal/services"
	repositories_mocks "prog/internal/repositories/mocks"
	managers_mocks "prog/internal/managers/mocks"
	"testing"
	"time"
	"context"
	"github.com/charmbracelet/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


type mockTrainingService struct {
	mockTrainingRepository     *repositories_mocks.MockTrainingRepository
	mockClientRepository       *repositories_mocks.MockClientRepository
	mockSubscriptionRepository *repositories_mocks.MockSubscriptionRepository
	mockHallRepository         *repositories_mocks.MockHallRepository
	mockTransactionManager     *managers_mocks.MockTransactionManager
	logger                     *log.Logger
}

func createMockTrainingService(controller *gomock.Controller) *mockTrainingService {
	service := new(mockTrainingService)

	service.mockTrainingRepository = repositories_mocks.NewMockTrainingRepository(controller)
	service.mockClientRepository = repositories_mocks.NewMockClientRepository(controller)
	service.mockSubscriptionRepository = repositories_mocks.NewMockSubscriptionRepository(controller)
	service.mockHallRepository = repositories_mocks.NewMockHallRepository(controller)
	service.mockTransactionManager = managers_mocks.NewMockTransactionManager(controller)
	service.logger = log.New(os.Stderr)

	return service
}

func createTrainingService(service *mockTrainingService) services.TrainingService {
	return NewTrainingServiceImplementation(service.mockTrainingRepository, service.mockClientRepository, service.mockSubscriptionRepository, service.mockHallRepository, service.mockTransactionManager, service.logger)
}

//-------------------------------------------------------------------------------------------------
// create

var testTrainingCreateSuccess = []struct {
	TestName  string
	InputData struct {
		training   *models.Training
	}
	Prepare     func(service *mockTrainingService)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple create",
		InputData: struct {
			training   *models.Training
		}{training: &models.Training{
					DateTime:  time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
					HallID:    7,
					CoachID:   10,
					PlacesNum: 13,
				}},

		Prepare: func(service *mockTrainingService) {
			ctx := context.Background()
			service.mockHallRepository.EXPECT().GetByID(ctx, uint64(7)).
				Return(&models.Hall{
							ID:       7,
							Capacity: 20,
						}, nil)

			service.mockTrainingRepository.EXPECT().GetAllByDateTime(ctx, time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC)).
				Return([]models.Training{
						{
							DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
							HallID:   111,
							CoachID:  13,
						},
						{
							DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
							HallID:   113,
							CoachID:  11,
						},
					}, nil)
			service.mockTrainingRepository.EXPECT().Create(ctx,
				&models.Training{
						DateTime:  time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
						HallID:    7,
						CoachID:   10,
						PlacesNum: 13,
					}).Return(nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
}


var testTrainingCreateFailure = []struct {
	TestName  string
	InputData struct {
		training   *models.Training
	}
	Prepare     func(service *mockTrainingService)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create error, not available time",
		InputData: struct {
			training   *models.Training
		}{training: &models.Training{
			DateTime:  time.Date(2024, 3, 5, 23, 0, 0, 0, time.UTC),
			HallID:    7,
			CoachID:   10,
			PlacesNum: 13,
		}},

		Prepare: func(service *mockTrainingService) {
		},

		CheckOutput: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, servicesErrors.IncorrectTrainingTime)
		},
	},
	{
		TestName: "create error, places number more then capacity",
		InputData: struct {
			training   *models.Training
		}{training: &models.Training{
			DateTime:  time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
			HallID:    7,
			CoachID:   10,
			PlacesNum: 13,
		}},

		Prepare: func(service *mockTrainingService) {
			ctx := context.Background()
			service.mockHallRepository.EXPECT().GetByID(ctx, uint64(7)).
			Return(&models.Hall{
						ID:       7,
						Capacity: 10,
					}, nil)
		},

		CheckOutput: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, servicesErrors.PlacesNumMoreThenCapacity)
		},
	},
	{
		TestName: "create error, hall busy on date time",
		InputData: struct {
			training   *models.Training
		}{training: &models.Training{
			DateTime:  time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
			HallID:    7,
			CoachID:   10,
			PlacesNum: 13,
		}},

		Prepare: func(service *mockTrainingService) {
			ctx := context.Background()
			service.mockHallRepository.EXPECT().GetByID(ctx, uint64(7)).
			Return(&models.Hall{
						ID:       7,
						Capacity: 20,
					}, nil)

			service.mockTrainingRepository.EXPECT().GetAllByDateTime(ctx, time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC)).
			Return([]models.Training{
					{
						DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
						HallID:   7,
						CoachID:  13,
					},
					{
						DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
						HallID:   113,
						CoachID:  11,
					},
				}, nil)
		},
		
		CheckOutput: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, servicesErrors.BysyDateTime)
		},
	},
	{
		TestName: "create error, coach busy on date time",
		InputData: struct {
			training   *models.Training
		}{training: &models.Training{
			DateTime:  time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
			HallID:    7,
			CoachID:   10,
			PlacesNum: 13,
		}},

		Prepare: func(service *mockTrainingService) {
			ctx := context.Background()
			service.mockHallRepository.EXPECT().GetByID(ctx, uint64(7)).
			Return(&models.Hall{
						ID:       7,
						Capacity: 20,
					}, nil)

			service.mockTrainingRepository.EXPECT().GetAllByDateTime(ctx, time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC)).
			Return([]models.Training{
					{
						DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
						HallID:   10,
						CoachID:  13,
					},
					{
						DateTime: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
						HallID:   113,
						CoachID:  10,
					},
				}, nil)
		},
		
		CheckOutput: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, servicesErrors.BysyDateTime)
		},
	},
}

func TestTrainingServiceImplementationCreate(t *testing.T) {
	for _, tt := range testTrainingCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := createMockTrainingService(ctrl)
			tt.Prepare(service)

			trainingService := createTrainingService(service)

			err := trainingService.Create(tt.InputData.training)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testTrainingCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := createMockTrainingService(ctrl)
			tt.Prepare(service)

			trainingService := createTrainingService(service)

			err := trainingService.Create(tt.InputData.training)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
