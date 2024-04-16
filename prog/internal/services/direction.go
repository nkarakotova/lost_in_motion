package services

import "prog/internal/models"

type DirectionService interface {
	Create(direction *models.Direction) error
	GetByID(id uint64) (*models.Direction, error)
	GetByName(name string) (*models.Direction, error)
	GetAll() ([]models.Direction, error)
}
