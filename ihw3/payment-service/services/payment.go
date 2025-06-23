package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"payment-service/infrastructure/kafka"
	"payment-service/infrastructure/storage"
	"payment-service/models"
	"time"
)

type PaymentWorker struct {
	accounts storage.AccountRepository
	inbox    kafka.Inboxer
	outbox   kafka.Outboxer
	manager  storage.Manager
	lg       *log.Logger
}

func NewPaymentWorker(accounts storage.AccountRepository, inbox kafka.Inboxer, outbox kafka.Outboxer, manager storage.Manager, lg *log.Logger) *PaymentWorker {
	return &PaymentWorker{
		accounts: accounts,
		inbox:    inbox,
		outbox:   outbox,
		manager:  manager,
		lg:       lg,
	}
}

func (s *PaymentWorker) Add(a *models.Account) error {
	return s.accounts.Add(a)
}

type paymentStatus struct {
	Id     uuid.UUID     `json:"id" db:"id"`
	Status models.Status `json:"status" db:"status"`
}

type paymentEvent struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Price  float64   `json:"price" db:"price"`
}

func (s *PaymentWorker) pay() (ok bool) {
	event, err := s.inbox.Get()
	if err != nil || event == nil {
		return false
	}
	pe := &paymentEvent{}
	err = json.Unmarshal(event.Payload, pe)

	if err != nil {
		s.lg.Printf("Error unmarshalling event: %s", err)
		return false
	}
	payment := models.NewPayment(event.ID, pe.UserID, pe.Price)
	//s.lg.Println("Payment: ", payment)

	t, err := s.manager.Begin()
	if err != nil {
		s.lg.Printf("Error begining transaction: %s", err)
		return false
	}
	defer t.Rollback()

	status := &paymentStatus{payment.ID, models.SUCCESS}

	err = s.accounts.PayWith(t, payment)
	if err != nil {
		status.Status = models.FAIL
		//s.lg.Println(err)
	}

	err = s.inbox.CompleteWith(t, event)
	if err != nil {
		s.lg.Println(err)
		return false
	}

	event, _ = models.NewEventWithJson(status)

	err = s.outbox.AddWith(t, event)
	if err != nil {
		s.lg.Println(err)
		return false
	}

	return t.Commit() == nil
}

func (s *PaymentWorker) StartPaying(ctx context.Context, period time.Duration) {
	ticker := time.NewTicker(period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if s.pay() {
					s.lg.Println("processed payment")
				}
			}
		}
	}()

}
