package kafka

import (
	"orders-service/models"
)

type Broker interface {
	Send(*models.Event) error
	Receive() (*models.Event, error)
	Close() error
	Register() error
}
