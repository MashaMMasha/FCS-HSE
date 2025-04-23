package services_test

import (
	"errors"
	"fmt"
	"testing"
	"zoo/application/services"
	"zoo/domain/animal"
	"zoo/domain/enclosure"
	"zoo/domain/events"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAnimalRepo struct {
	mock.Mock
	lastID int
}

func (m *MockAnimalRepo) GetByID(id int) (*animal.Animal, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*animal.Animal), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAnimalRepo) Save(animal *animal.Animal) error {
	if animal.ID == 0 {
		m.lastID++
		animal.ID = m.lastID
	}
	args := m.Called(animal)
	return args.Error(0)
}
func (m *MockAnimalRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAnimalRepo) GetAll() ([]*animal.Animal, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*animal.Animal), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockEnclosureRepo struct {
	mock.Mock
}

func (m *MockEnclosureRepo) GetByID(id int) (*enclosure.Enclosure, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*enclosure.Enclosure), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockEnclosureRepo) Save(enclosure *enclosure.Enclosure) error {
	args := m.Called(enclosure)
	return args.Error(0)
}

func (m *MockEnclosureRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEnclosureRepo) GetAll() []*enclosure.Enclosure {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*enclosure.Enclosure)
	}
	return nil
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) HandleAnimalMoved(event events.AnimalMovedEvent) {
	m.Called(event)
}

func (m *MockEventHandler) HandleFeedingTime(event events.FeedingTimeEvent) {
	m.Called(event)
}
func TestTransferAnimal_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)
	eventHandler := new(MockEventHandler)

	animalEntity := &animal.Animal{ID: 1, EnclosureID: 100}
	from := &enclosure.Enclosure{ID: 100, AnimalIDs: map[int]struct{}{1: {}}, MaxCapacity: 10}
	to := &enclosure.Enclosure{ID: 200, AnimalIDs: map[int]struct{}{}, MaxCapacity: 10}

	animalRepo.On("GetByID", 1).Return(animalEntity, nil)
	enclosureRepo.On("GetByID", 100).Return(from, nil)
	enclosureRepo.On("GetByID", 200).Return(to, nil)
	enclosureRepo.On("Save", from).Return(nil)
	enclosureRepo.On("Save", to).Return(nil)
	animalRepo.On("Save", animalEntity).Return(nil)
	eventHandler.On("HandleAnimalMoved", mock.AnythingOfType("events.AnimalMovedEvent")).Return()

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventHandler)

	event, err := service.TransferAnimal(1, 200)
	assert.NoError(t, err)
	assert.Equal(t, 1, event.AnimalID)
	assert.Equal(t, 100, event.OldEnclosureID)
	assert.Equal(t, 200, event.NewEnclosureID)

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}

func TestTransferAnimal_AnimalNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)
	eventHandler := new(MockEventHandler)

	animalRepo.On("GetByID", 1).Return(nil, errors.New("animal not found"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventHandler)

	_, err := service.TransferAnimal(1, 200)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "animal not found")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}
func TestAddAnimal_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	animalRepo.lastID = 0
	enclosureRepo := new(MockEnclosureRepo)

	enclosureEntity := &enclosure.Enclosure{ID: 100, MaxCapacity: 10, AnimalIDs: map[int]struct{}{}, AnimalType: "predator"}

	enclosureRepo.On("GetByID", 100).Return(enclosureEntity, nil)
	animalRepo.On("Save", mock.AnythingOfType("*animal.Animal")).Return(nil)
	enclosureRepo.On("Save", enclosureEntity).Return(nil)

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)

	animalID, err := service.AddAnimal(100, "Tralalelo tralala", "Bombardiro crocodilo", animal.Predator)
	assert.NoError(t, err)
	assert.Equal(t, 1, animalID)
	fmt.Println(animalID)
	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
}
func TestDeleteAnimal_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	animalEntity := &animal.Animal{ID: 1, EnclosureID: 100}
	enclosureEntity := &enclosure.Enclosure{ID: 100, AnimalIDs: map[int]struct{}{1: {}}}

	animalRepo.On("GetByID", 1).Return(animalEntity, nil)
	enclosureRepo.On("GetByID", 100).Return(enclosureEntity, nil)
	enclosureRepo.On("Save", enclosureEntity).Return(nil)
	animalRepo.On("Delete", 1).Return(nil)

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)

	err := service.DeleteAnimal(1)
	assert.NoError(t, err)

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
}

func TestAddEnclosure_Success(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureRepo.On("Save", mock.AnythingOfType("*enclosure.Enclosure")).Return(nil)

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	enclosureID, err := service.AddEnclosure(animal.Predator, 10)
	assert.NoError(t, err)
	assert.Equal(t, 0, enclosureID)

	enclosureRepo.AssertExpectations(t)
}

func TestDeleteEnclosure_Success(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureEntity := &enclosure.Enclosure{ID: 1, AnimalCount: 0}
	enclosureRepo.On("GetByID", 1).Return(enclosureEntity, nil)
	enclosureRepo.On("Delete", 1).Return(nil)

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	err := service.DeleteEnclosure(1)
	assert.NoError(t, err)

	enclosureRepo.AssertExpectations(t)
}

func TestDeleteEnclosure_NotEmpty(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureEntity := &enclosure.Enclosure{ID: 1, AnimalCount: 5}
	enclosureRepo.On("GetByID", 1).Return(enclosureEntity, nil)

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	err := service.DeleteEnclosure(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enclosure is not empty")

	enclosureRepo.AssertExpectations(t)
}
func TestGetAnimalByID_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)

	animalEntity := &animal.Animal{ID: 1, Name: "Lion"}
	animalRepo.On("GetByID", 1).Return(animalEntity, nil)

	service := services.NewAnimalTransferService(animalRepo, nil, nil)

	result, err := service.GetAnimalByID(1)
	assert.NoError(t, err)
	assert.Equal(t, animalEntity, result)

	animalRepo.AssertExpectations(t)
}

func TestGetAnimalByID_NotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)

	animalRepo.On("GetByID", 1).Return(nil, errors.New("animal not found"))

	service := services.NewAnimalTransferService(animalRepo, nil, nil)

	result, err := service.GetAnimalByID(1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "animal not found")

	animalRepo.AssertExpectations(t)
}

func TestGetAllAnimals_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)

	animals := []*animal.Animal{
		{ID: 1, Name: "Lion"},
		{ID: 2, Name: "Tiger"},
	}
	animalRepo.On("GetAll").Return(animals, nil)

	service := services.NewAnimalTransferService(animalRepo, nil, nil)

	result, err := service.GetAllAnimals()
	assert.NoError(t, err)
	assert.Equal(t, animals, result)

	animalRepo.AssertExpectations(t)
}
func TestGetAllEnclosures_Success(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosures := []*enclosure.Enclosure{
		{ID: 1, AnimalType: "predator", MaxCapacity: 10},
		{ID: 2, AnimalType: "predator", MaxCapacity: 10},
	}
	enclosureRepo.On("GetAll").Return(enclosures)

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	result, err := service.GetAllEnclosures()
	assert.NoError(t, err)
	assert.Equal(t, enclosures, result)

	enclosureRepo.AssertExpectations(t)
}

func TestGetEnclosureByID_Success(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureEntity := &enclosure.Enclosure{ID: 1, AnimalType: "predator", MaxCapacity: 10}
	enclosureRepo.On("GetByID", 1).Return(enclosureEntity, nil)

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	result, err := service.GetEnclosureByID(1)
	assert.NoError(t, err)
	assert.Equal(t, enclosureEntity, result)

	enclosureRepo.AssertExpectations(t)
}

func TestGetEnclosureByID_NotFound(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureRepo.On("GetByID", 1).Return(nil, errors.New("enclosure not found"))

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	result, err := service.GetEnclosureByID(1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enclosure not found")

	enclosureRepo.AssertExpectations(t)
}

func TestTransferAnimal_SourceEnclosureNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)
	eventHandler := new(MockEventHandler)

	animalEntity := &animal.Animal{ID: 1, EnclosureID: 100}
	animalRepo.On("GetByID", 1).Return(animalEntity, nil)
	enclosureRepo.On("GetByID", 100).Return(nil, errors.New("source enclosure not found"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventHandler)

	_, err := service.TransferAnimal(1, 200)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "first enclosure not found")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}

func TestTransferAnimal_DestinationEnclosureNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)
	eventHandler := new(MockEventHandler)

	animalEntity := &animal.Animal{ID: 1, EnclosureID: 100}
	from := &enclosure.Enclosure{ID: 100, AnimalIDs: map[int]struct{}{1: {}}}
	animalRepo.On("GetByID", 1).Return(animalEntity, nil)
	enclosureRepo.On("GetByID", 100).Return(from, nil)
	enclosureRepo.On("GetByID", 200).Return(nil, errors.New("destination enclosure not found"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventHandler)

	_, err := service.TransferAnimal(1, 200)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "second enclosure not found")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}

func TestAddAnimal_EnclosureNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	enclosureRepo.On("GetByID", 100).Return(nil, errors.New("enclosure not found"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)

	_, err := service.AddAnimal(100, "Meat", "Lion", animal.Predator)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enclosure not found")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
}

func TestDeleteAnimal_AnimalNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	animalRepo.On("GetByID", 1).Return(nil, errors.New("animal not found"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)

	err := service.DeleteAnimal(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "animal not found")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
}

func TestDeleteEnclosure_EnclosureNotFound(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureRepo.On("GetByID", 1).Return(nil, errors.New("enclosure not found"))

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	err := service.DeleteEnclosure(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enclosure not found")

	enclosureRepo.AssertExpectations(t)
}

func TestDeleteAnimal_SaveEnclosureError(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	enclosureRepo := new(MockEnclosureRepo)

	animalEntity := &animal.Animal{ID: 1, EnclosureID: 100}
	enclosureEntity := &enclosure.Enclosure{ID: 100, AnimalIDs: map[int]struct{}{1: {}}}

	animalRepo.On("GetByID", 1).Return(animalEntity, nil)
	enclosureRepo.On("GetByID", 100).Return(enclosureEntity, nil)
	enclosureRepo.On("Save", enclosureEntity).Return(errors.New("failed to update enclosure"))

	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)

	err := service.DeleteAnimal(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update enclosure")

	animalRepo.AssertExpectations(t)
	enclosureRepo.AssertExpectations(t)
}

func TestDeleteEnclosure_DeleteError(t *testing.T) {
	enclosureRepo := new(MockEnclosureRepo)

	enclosureEntity := &enclosure.Enclosure{ID: 1, AnimalCount: 0}
	enclosureRepo.On("GetByID", 1).Return(enclosureEntity, nil)
	enclosureRepo.On("Delete", 1).Return(errors.New("failed to delete enclosure"))

	service := services.NewAnimalTransferService(nil, enclosureRepo, nil)

	err := service.DeleteEnclosure(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete enclosure")

	enclosureRepo.AssertExpectations(t)
}
