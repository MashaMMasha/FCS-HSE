package repositories

import (
	"errors"
	"sync"
	"zoo/domain/enclosure"
)

type InMemoryEnclosureRepository struct {
	data   map[int]*enclosure.Enclosure
	nextID int
	mutex  sync.Mutex
}

func NewInMemoryEnclosureRepository() *InMemoryEnclosureRepository {
	return &InMemoryEnclosureRepository{
		data:   make(map[int]*enclosure.Enclosure),
		nextID: 1,
	}
}

func (repo *InMemoryEnclosureRepository) GetByID(id int) (*enclosure.Enclosure, error) {
	if a, ok := repo.data[id]; ok {
		return a, nil
	}
	return nil, errors.New("enclosure not found")
}

func (repo *InMemoryEnclosureRepository) Save(e *enclosure.Enclosure) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if e.ID == 0 {
		e.ID = repo.nextID
		repo.nextID++
	}
	repo.data[e.ID] = e
	return nil
}

func (repo *InMemoryEnclosureRepository) Delete(id int) error {
	delete(repo.data, id)
	return nil
}

func (repo *InMemoryEnclosureRepository) GetAll() []*enclosure.Enclosure {
	var result []*enclosure.Enclosure
	for _, a := range repo.data {
		result = append(result, a)
	}
	return result
}
