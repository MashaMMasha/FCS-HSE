package services

import (
	"fmt"
	"time"
	"zoo/application/ports"
	"zoo/domain/events"
	"zoo/domain/schedule"
)

type FeedingService struct {
	animalRepo   ports.AnimalRepository
	scheduleRepo ports.ScheduleRepository
	eventHandler ports.EventHandler
}

func NewFeedingService(
	animalRepo ports.AnimalRepository,
	scheduleRepo ports.ScheduleRepository,
	eventHandler ports.EventHandler,
) *FeedingService {
	return &FeedingService{
		animalRepo:   animalRepo,
		scheduleRepo: scheduleRepo,
		eventHandler: eventHandler,
	}
}

func (s *FeedingService) FeedAnimal(animalID int, food string) error {
	animal, err := s.animalRepo.GetByID(animalID)
	if err != nil {
		return fmt.Errorf("failed to get animal: %w", err)
	}

	schedule, err := s.scheduleRepo.GetByAnimalID(animalID)
	if err != nil {
		return fmt.Errorf("failed to get feeding schedule: %w", err)
	}

	if err = animal.Feed(food); err != nil {
		return fmt.Errorf("failed to feed animal: %w", err)
	}

	schedule.LastFedTime = time.Now()
	if err = s.scheduleRepo.Save(schedule); err != nil {
		return fmt.Errorf("failed to update schedule: %w", err)
	}

	if err = s.animalRepo.Save(animal); err != nil {
		return fmt.Errorf("failed to save animal: %w", err)
	}

	event := events.FeedingTimeEvent{
		AnimalID:    animalID,
		ScheduleID:  schedule.ID,
		FoodType:    food,
		FeedingTime: time.Now(),
	}

	if s.eventHandler != nil {
		s.eventHandler.HandleFeedingTime(event)
	}

	return nil
}

func (s *FeedingService) AddFeedingSchedule(animalID int, foodType string, feedingInterval int) error {
	animal, err := s.animalRepo.GetByID(animalID)
	if err != nil {
		return fmt.Errorf("failed to get animal: %w", err)
	}

	schedule := schedule.FeedingSchedule{
		AnimalID:     animal.ID,
		FoodType:     foodType,
		FeedInterval: feedingInterval,
		LastFedTime:  time.Time{},
	}

	if err = s.scheduleRepo.Save(&schedule); err != nil {
		return fmt.Errorf("failed to save feeding schedule: %w", err)
	}

	return nil
}

func (s *FeedingService) ChangeFeedInterval(scheduleID int, newInterval int) error {
	schedule, err := s.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return fmt.Errorf("failed to get feeding schedule: %w", err)
	}

	schedule.ChangeFeedInterval(newInterval)

	if err = s.scheduleRepo.Save(schedule); err != nil {
		return fmt.Errorf("failed to update feeding schedule: %w", err)
	}

	return nil
}
