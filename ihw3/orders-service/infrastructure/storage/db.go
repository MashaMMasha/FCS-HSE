package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"orders-service/models"
)

type OrderDB struct {
	db *sqlx.DB
}

const (
	createQuery = `
		CREATE TABLE IF NOT EXISTS orders (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			status TEXT NOT NULL,
			price NUMERIC NOT NULL,
			descr TEXT
		);		
	`
	addQuery          = `INSERT INTO orders (id, user_id, status, price, descr) VALUES (:id, :user_id, :status, :price, :descr)`
	updateStatusQuery = `UPDATE orders SET status=$2 WHERE id=$1`
	allQuery          = `SELECT * FROM orders`
	getQuery          = `SELECT * FROM orders WHERE id=$1`
)

func NewOrderDB(db *sqlx.DB) (*OrderDB, error) {
	_, err := db.Exec(createQuery)

	if err != nil {
		return nil, err
	}

	return &OrderDB{db}, nil
}

func (odb *OrderDB) Get(id uuid.UUID) (order *models.Order, err error) {
	order = &models.Order{}
	err = odb.db.Get(order, getQuery, id)

	if err != nil {
		return nil, err
	}

	return
}

func (odb *OrderDB) AddWith(transaction Transaction, order *models.Order) (err error) {
	tx, ok := transaction.(*sqlx.Tx)
	if !ok {
		return fmt.Errorf("transaction is not a sqlx.Tx")
	}

	_, err = tx.NamedExec(addQuery, order)

	return
}
func (odb *OrderDB) Add(order *models.Order) (err error) {
	_, err = odb.db.NamedExec(addQuery, order)

	return
}
func (odb *OrderDB) All() (f []*models.Order, err error) {
	err = odb.db.Select(&f, allQuery)

	if err != nil {
		return nil, err
	}

	return
}

func (odb *OrderDB) UpdateStatus(id uuid.UUID, status models.Status) (err error) {
	res, err := odb.db.Exec(updateStatusQuery, id, status)
	log.Println(res)
	return
}
