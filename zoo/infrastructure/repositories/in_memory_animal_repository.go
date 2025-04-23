package repositories

import (
	"errors"
	"sync"
	"zoo/domain/animal"
)

type InMemoryAnimalRepository struct {
	data   map[int]*animal.Animal
	nextID int
	mutex  sync.Mutex
}

func NewInMemoryAnimalRepository() *InMemoryAnimalRepository {
	return &InMemoryAnimalRepository{
		data:   make(map[int]*animal.Animal),
		nextID: 1,
	}
}

func (repo *InMemoryAnimalRepository) GetByID(id int) (*animal.Animal, error) {
	if a, ok := repo.data[id]; ok {
		return a, nil
	}
	return nil, errors.New("animal not found")
}

func (repo *InMemoryAnimalRepository) Save(a *animal.Animal) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if a.ID == 0 {
		a.ID = repo.nextID
		repo.nextID++
	}
	repo.data[a.ID] = a
	return nil
}

func (repo *InMemoryAnimalRepository) Delete(id int) error {
	delete(repo.data, id)
	return nil
}

func (repo *InMemoryAnimalRepository) GetAll() ([]*animal.Animal, error) {
	var result []*animal.Animal
	for _, a := range repo.data {
		result = append(result, a)
	}
	return result, nil
}
