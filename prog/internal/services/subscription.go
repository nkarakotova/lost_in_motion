package services

import "prog/internal/models"

type SubscriptionService interface {
	Create(subscription *models.Subscription, clientID uint64) error
	GetByID(id uint64) (*models.Subscription, error)
}
