package storage

import (
	"github.com/google/uuid"
	"orders-service/models"
)

type OrdersRepository interface {
	Add(*models.Order) error
	AddWith(Transaction, *models.Order) error
	Get(uuid.UUID) (*models.Order, error)
	All() ([]*models.Order, error)
	UpdateStatus(uuid.UUID, models.Status) error
}
