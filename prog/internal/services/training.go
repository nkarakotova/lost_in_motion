package services

import (
	"prog/internal/models"
	"time"
)

type TrainingService interface {
	Create(training *models.Training) error
	Delete(id uint64) error
	GetByID(id uint64) (*models.Training, error)
	GetAllByClient(id uint64) ([]models.Training, error)
	GetAllByCoachOnDate(id uint64, date time.Time) ([]models.Training, error)
	GetAllByDateTime(dateTime time.Time) ([]models.Training, error)
	GetAllBetweenDateTime(start time.Time, end time.Time) ([]models.Training, error)
}

const FirstTrainingTime = 10
const LastTrainingTime = 22
