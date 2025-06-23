package kafka

import (
	"payment-service/models"
)

type Broker interface {
	Send(*models.Event) error
	Receive() (*models.Event, error)
	Close() error
	Register() error
}
