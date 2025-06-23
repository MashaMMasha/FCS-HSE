package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"orders-service/infrastructure/kafka"
	"orders-service/infrastructure/storage"
	"orders-service/models"
)

type OrderServicing interface {
	Add(string, float64, string) (string, error)
	Get(userID string) (*models.Order, error)
	All() ([]*models.Order, error)
	UpdateStatus(string, string) error
}

type OrderService struct {
	storage storage.OrdersRepository
	outbox  kafka.Outboxer
	manager storage.Manager
}

func NewOrderService(storage storage.OrdersRepository, outbox kafka.Outboxer, manager storage.Manager) *OrderService {
	return &OrderService{storage: storage, outbox: outbox, manager: manager}
}

type orderPayload struct {
	//ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Price  float64   `json:"price"`
}

func (s *OrderService) newOrderEvent(order *models.Order) *models.Event {
	op := orderPayload{
		UserID: order.UserID,
		Price:  order.Price,
	}

	payload, _ := json.Marshal(op)
	return models.NewEventWithID(order.ID, payload)
}

func (s *OrderService) Add(userID string, price float64, descr string) (string, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return "", err
	}

	if price < 0 {
		return "", fmt.Errorf("price must be greater than zero")
	}

	order := models.NewOrder(uid, price, descr)

	tx, err := s.manager.Begin()

	defer tx.Rollback()
	if err != nil {
		return "", err
	}

	err = s.storage.AddWith(tx, order)
	if err != nil {
		return "", err
	}

	event := s.newOrderEvent(order)

	err = s.outbox.AddWith(tx, event)
	if err != nil {
		return "", err
	}

	return order.ID.String(), tx.Commit()
}

func (s *OrderService) Get(userID string) (order *models.Order, err error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return
	}
	order, err = s.storage.Get(uid)
	return
}

func (s *OrderService) All() ([]*models.Order, error) {
	return s.storage.All()
}

func (s *OrderService) UpdateStatus(id string, status string) (err error) {
	sstatus, err := models.ParseStatus(status)
	if err != nil {
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}

	return s.storage.UpdateStatus(uid, sstatus)
}
