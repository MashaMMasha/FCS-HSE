package services_test

import (
	"errors"
	"testing"
	"zoo/application/services"
	"zoo/domain/animal"
	"zoo/domain/schedule"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockScheduleRepo struct {
	mock.Mock
}

func (m *MockScheduleRepo) GetByID(id int) (*schedule.FeedingSchedule, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*schedule.FeedingSchedule), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockScheduleRepo) GetByAnimalID(animalID int) (*schedule.FeedingSchedule, error) {
	args := m.Called(animalID)
	if args.Get(0) != nil {
		return args.Get(0).(*schedule.FeedingSchedule), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockScheduleRepo) Save(schedule *schedule.FeedingSchedule) error {
	args := m.Called(schedule)
	return args.Error(0)
}

func (m *MockScheduleRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestFeedAnimal_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	scheduleRepo := new(MockScheduleRepo)
	eventHandler := new(MockEventHandler)

	animalObj := &animal.Animal{ID: 1, FavoriteFood: "meat"}
	scheduleObj := &schedule.FeedingSchedule{ID: 2, AnimalID: 1}

	animalRepo.On("GetByID", 1).Return(animalObj, nil)
	scheduleRepo.On("GetByAnimalID", 1).Return(scheduleObj, nil)
	animalRepo.On("Save", animalObj).Return(nil)
	scheduleRepo.On("Save", scheduleObj).Return(nil)
	eventHandler.On("HandleFeedingTime", mock.AnythingOfType("events.FeedingTimeEvent")).Return()

	service := services.NewFeedingService(animalRepo, scheduleRepo, eventHandler)

	err := service.FeedAnimal(1, "meat")

	assert.NoError(t, err)
	animalRepo.AssertExpectations(t)
	scheduleRepo.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}

func TestFeedAnimal_AnimalNotFound(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	scheduleRepo := new(MockScheduleRepo)

	animalRepo.On("GetByID", 1).Return(nil, errors.New("not found"))

	service := services.NewFeedingService(animalRepo, scheduleRepo, nil)

	err := service.FeedAnimal(1, "meat")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get animal")
}

func TestAddFeedingSchedule_Success(t *testing.T) {
	animalRepo := new(MockAnimalRepo)
	scheduleRepo := new(MockScheduleRepo)

	animalObj := &animal.Animal{ID: 1}
	animalRepo.On("GetByID", 1).Return(animalObj, nil)
	scheduleRepo.On("Save", mock.AnythingOfType("*schedule.FeedingSchedule")).Return(nil)

	service := services.NewFeedingService(animalRepo, scheduleRepo, nil)

	err := service.AddFeedingSchedule(1, "grass", 6)

	assert.NoError(t, err)
	animalRepo.AssertExpectations(t)
	scheduleRepo.AssertExpectations(t)
}

func TestChangeFeedInterval_Success(t *testing.T) {
	scheduleRepo := new(MockScheduleRepo)

	scheduleObj := &schedule.FeedingSchedule{ID: 3, FeedInterval: 4}
	scheduleRepo.On("GetByID", 3).Return(scheduleObj, nil)
	scheduleRepo.On("Save", scheduleObj).Return(nil)

	service := services.NewFeedingService(nil, scheduleRepo, nil)

	err := service.ChangeFeedInterval(3, 8)

	assert.NoError(t, err)
	assert.Equal(t, 8, scheduleObj.FeedInterval)
	scheduleRepo.AssertExpectations(t)
}
