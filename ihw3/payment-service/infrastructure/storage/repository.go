package storage

import (
	"github.com/google/uuid"
	"payment-service/models"
)

type AccountRepository interface {
	Add(*models.Account) error
	Get(uuid.UUID) (*models.Account, error)
	All() ([]*models.Account, error)
	Update(id uuid.UUID, amount float64) (err error)
	PayWith(Transaction, *models.Payment) error
}
