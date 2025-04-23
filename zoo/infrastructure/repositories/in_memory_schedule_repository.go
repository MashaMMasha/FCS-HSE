package repositories

import (
	"errors"
	"zoo/domain/schedule"
)

type InMemoryScheduleRepository struct {
	data map[int]*schedule.FeedingSchedule
}

func (repo *InMemoryScheduleRepository) GetByAnimalID(animalID int) (*schedule.FeedingSchedule, error) {
	for i := range repo.data {
		if repo.data[i].AnimalID == animalID {
			return repo.data[i], nil
		}
	}
	return nil, errors.New("schedule not found")
}

func NewInMemoryScheduleRepository() *InMemoryScheduleRepository {
	return &InMemoryScheduleRepository{
		data: make(map[int]*schedule.FeedingSchedule),
	}
}

func (repo *InMemoryScheduleRepository) GetByID(id int) (*schedule.FeedingSchedule, error) {
	if a, ok := repo.data[id]; ok {
		return a, nil
	}
	return nil, errors.New("schedule not found")
}

func (repo *InMemoryScheduleRepository) Save(a *schedule.FeedingSchedule) error {
	repo.data[a.ID] = a
	return nil
}

func (repo *InMemoryScheduleRepository) Delete(id int) error {
	delete(repo.data, id)
	return nil
}

func (repo *InMemoryScheduleRepository) GetAll() []*schedule.FeedingSchedule {
	var result []*schedule.FeedingSchedule
	for _, a := range repo.data {
		result = append(result, a)
	}
	return result
}
