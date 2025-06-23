package services

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"payment-service/infrastructure/storage"
	"payment-service/models"
)

type AccountServicing interface {
	Add(string, float64) (string, error)
	Get(id string) (*models.Account, error)
	All() ([]*models.Account, error)
	Update(string, float64) error
}

type AccountService struct {
	accounts storage.AccountRepository
	lg       *log.Logger
}

func NewAccountService(accounts storage.AccountRepository, lg *log.Logger) *AccountService {
	return &AccountService{accounts: accounts, lg: lg}
}

func (s *AccountService) Add(fullname string, balance float64) (string, error) {
	uid, err := uuid.NewRandom()
	id := uid.String()
	if err != nil {
		return "", err
	}
	if balance < 0 {
		return "", errors.New("balance must be greater than zero")
	}
	if fullname == "" {
		return "", errors.New("full name must not be empty")
	}

	return id, s.accounts.Add(models.NewAccount(uid, fullname, balance))
}

func (s *AccountService) Get(userID string) (*models.Account, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.accounts.Get(uid)
}

func (s *AccountService) All() ([]*models.Account, error) {
	return s.accounts.All()
}

func (s *AccountService) Update(userID string, amount float64) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	return s.accounts.Update(uid, amount)
}
