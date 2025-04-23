package services_test

import (
	"errors"
	"testing"
	"zoo/application/services"
	"zoo/domain/animal"
	"zoo/domain/enclosure"

	"github.com/stretchr/testify/assert"
)

func TestGetAnimalCount_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	animals := []*animal.Animal{
		{ID: 1, Name: "Lion"},
		{ID: 2, Name: "Tiger"},
	}
	animalRepo.On("GetAll").Return(animals, nil)

	service := services.NewZooStatisticsService(animalRepo, enclosureRepo)

	count, err := service.GetAnimalCount()

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	animalRepo.AssertExpectations(t)
}

func TestGetAnimalCount_Error(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	animalRepo.On("GetAll").Return(nil, errors.New("database error"))

	service := services.NewZooStatisticsService(animalRepo, enclosureRepo)

	count, err := service.GetAnimalCount()

	assert.Error(t, err)
	assert.Equal(t, 0, count)
	animalRepo.AssertExpectations(t)
}

func TestGetEnclosureCount_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	enclosures := []*enclosure.Enclosure{
		{ID: 1, AnimalType: "predator", MaxCapacity: 10},
		{ID: 2, AnimalType: "herbivore", MaxCapacity: 20},
	}
	enclosureRepo.On("GetAll").Return(enclosures)

	service := services.NewZooStatisticsService(animalRepo, enclosureRepo)

	count, err := service.GetEnclosureCount()

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	enclosureRepo.AssertExpectations(t)
}
