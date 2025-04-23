package events

import "time"

type FeedingTimeEvent struct {
	AnimalID    int
	ScheduleID  int
	FoodType    string
	FeedingTime time.Time
}

func NewFeedingTimeEvent(animalID, scheduleID int, foodType string) FeedingTimeEvent {
	return FeedingTimeEvent{
		AnimalID:    animalID,
		ScheduleID:  scheduleID,
		FoodType:    foodType,
		FeedingTime: time.Now(),
	}
}
