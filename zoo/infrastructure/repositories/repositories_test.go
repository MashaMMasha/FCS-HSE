package repositories_test

import (
	"testing"
	"zoo/domain/animal"
	"zoo/domain/enclosure"
	"zoo/domain/schedule"
	"zoo/infrastructure/repositories"

	"github.com/stretchr/testify/assert"
)

func TestAnimalRepository_Save_Success(t *testing.T) {
	repo := repositories.NewInMemoryAnimalRepository()

	animalObj := &animal.Animal{ID: 1, Name: "Lion"}
	err := repo.Save(animalObj)

	assert.NoError(t, err)

	savedAnimal, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, animalObj, savedAnimal)
}

func TestAnimalRepository_GetByID_NotFound(t *testing.T) {
	repo := repositories.NewInMemoryAnimalRepository()

	_, err := repo.GetByID(1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestAnimalRepository_Delete_Success(t *testing.T) {
	repo := repositories.NewInMemoryAnimalRepository()

	animalObj := &animal.Animal{ID: 1, Name: "Lion"}
	_ = repo.Save(animalObj)

	err := repo.Delete(1)
	assert.NoError(t, err)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestAnimalRepository_GetAll_Success(t *testing.T) {
	repo := repositories.NewInMemoryAnimalRepository()

	animal1 := &animal.Animal{ID: 0, Name: "Lion"}
	animal2 := &animal.Animal{ID: 2, Name: "Tiger"}
	_ = repo.Save(animal1)
	_ = repo.Save(animal2)

	animals, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, animals, 2)
	assert.Contains(t, animals, animal1)
	assert.Contains(t, animals, animal2)
}

func TestEnclosureRepository_Save_Success(t *testing.T) {
	repo := repositories.NewInMemoryEnclosureRepository()

	enclosure := &enclosure.Enclosure{ID: 0, AnimalType: "predator", MaxCapacity: 10}
	err := repo.Save(enclosure)

	assert.NoError(t, err)

	savedEnclosure, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, enclosure, savedEnclosure)
}

func TestEnclosureRepository_GetByID_NotFound(t *testing.T) {
	repo := repositories.NewInMemoryEnclosureRepository()

	_, err := repo.GetByID(1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestEnclosureRepository_Delete_Success(t *testing.T) {
	repo := repositories.NewInMemoryEnclosureRepository()

	enclosure := &enclosure.Enclosure{ID: 0, AnimalType: "predator", MaxCapacity: 10}
	_ = repo.Save(enclosure)

	err := repo.Delete(1)
	assert.NoError(t, err)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestEnclosureRepository_GetAll_Success(t *testing.T) {
	repo := repositories.NewInMemoryEnclosureRepository()

	enclosure1 := &enclosure.Enclosure{ID: 1, AnimalType: "predator", MaxCapacity: 10}
	enclosure2 := &enclosure.Enclosure{ID: 2, AnimalType: "predator", MaxCapacity: 10}
	_ = repo.Save(enclosure1)
	_ = repo.Save(enclosure2)

	enclosures := repo.GetAll()
	assert.Len(t, enclosures, 2)
	assert.Contains(t, enclosures, enclosure1)
	assert.Contains(t, enclosures, enclosure2)
}

func TestInMemoryScheduleRepository_Save(t *testing.T) {
	repo := repositories.NewInMemoryScheduleRepository()
	schedule := &schedule.FeedingSchedule{ID: 1, AnimalID: 1}
	err := repo.Save(schedule)
	assert.NoError(t, err)

	savedSchedule, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, schedule, savedSchedule)
}

func TestInMemoryScheduleRepository_GetByID_NotFound(t *testing.T) {
	repo := repositories.NewInMemoryScheduleRepository()
	_, err := repo.GetByID(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "schedule not found")
}

func TestInMemoryScheduleRepository_Delete(t *testing.T) {
	repo := repositories.NewInMemoryScheduleRepository()
	schedule := &schedule.FeedingSchedule{ID: 1, AnimalID: 1}
	_ = repo.Save(schedule)

	err := repo.Delete(1)
	assert.NoError(t, err)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "schedule not found")
}

func TestInMemoryScheduleRepository_GetAll(t *testing.T) {
	repo := repositories.NewInMemoryScheduleRepository()
	schedule1 := &schedule.FeedingSchedule{ID: 1, AnimalID: 1}
	schedule2 := &schedule.FeedingSchedule{ID: 2, AnimalID: 2}
	_ = repo.Save(schedule1)
	_ = repo.Save(schedule2)

	schedules := repo.GetAll()
	assert.Len(t, schedules, 2)
	assert.Contains(t, schedules, schedule1)
	assert.Contains(t, schedules, schedule2)
}

func TestInMemoryScheduleRepository_GetByAnimalID(t *testing.T) {
	repo := repositories.NewInMemoryScheduleRepository()
	schedule1 := &schedule.FeedingSchedule{ID: 1, AnimalID: 1}
	schedule2 := &schedule.FeedingSchedule{ID: 2, AnimalID: 2}
	_ = repo.Save(schedule1)
	_ = repo.Save(schedule2)

	schedule, err := repo.GetByAnimalID(1)
	assert.NoError(t, err)
	assert.Equal(t, schedule1, schedule)

	_, err = repo.GetByAnimalID(3)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "schedule not found")
}
