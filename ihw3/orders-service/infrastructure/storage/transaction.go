package storage

import "github.com/jmoiron/sqlx"

type Transaction interface {
	Commit() error
	Rollback() error
}

type Manager interface {
	Begin() (Transaction, error)
}

type DBManager struct {
	db *sqlx.DB
}

func NewDBManager(db *sqlx.DB) *DBManager {
	return &DBManager{db: db}
}

func (m *DBManager) Begin() (Transaction, error) {
	return m.db.Beginx()
}
