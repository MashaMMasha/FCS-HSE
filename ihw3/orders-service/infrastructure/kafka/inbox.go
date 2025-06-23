package kafka

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"orders-service/infrastructure/storage"
	"orders-service/models"
)

type Inboxer interface {
	Add(*models.Event) error
	Get() (*models.Event, error)
	CompleteWith(storage.Transaction, *models.Event) error
}

type Inbox struct {
	db *sqlx.DB
}

func NewInbox(db *sqlx.DB) (*Inbox, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS inbox (
			id UUID PRIMARY KEY,
			created_at TIMESTAMP,
			processed BOOLEAN NOT NULL,
			type TEXT NOT NULL,
			payload JSONB
		);
		
	`)

	if err != nil {
		return nil, err
	}

	return &Inbox{db}, nil
}

func (i *Inbox) Add(e *models.Event) (err error) {
	_, err = i.db.NamedExec(`INSERT INTO inbox (id, created_at, processed, type, payload) VALUES (:id, :created_at, :processed, :type, :payload)`, e)

	return
}
func (i *Inbox) Get() (event *models.Event, err error) {
	event = &models.Event{}
	err = i.db.Get(event, `SELECT * FROM inbox WHERE processed = false LIMIT 1`)
	return
}

func (i *Inbox) CompleteWith(transaction storage.Transaction, event *models.Event) (err error) {
	tx, ok := transaction.(*sqlx.Tx)

	if !ok {
		return fmt.Errorf(`trx is not a sqlx.Tx`)
	}

	_, err = tx.NamedExec(`UPDATE inbox SET processed = true WHERE id = :id`, event)
	return
}
