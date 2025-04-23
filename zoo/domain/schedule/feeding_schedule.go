package schedule

import (
	"fmt"
	"time"
)

type FeedingSchedule struct {
	ID           int
	AnimalID     int
	FeedInterval int
	LastFedTime  time.Time
	FoodType     string
}

func NewFeedingSchedule(animalID, feedInterval int, foodType string) *FeedingSchedule {
	return &FeedingSchedule{
		AnimalID:     animalID,
		FeedInterval: feedInterval,
		LastFedTime:  time.Now(),
		FoodType:     foodType,
	}
}

func (fs *FeedingSchedule) FeedAnimal() error {
	nextFeedingTime := fs.LastFedTime.Add(time.Duration(fs.FeedInterval) * time.Hour)

	if time.Now().Before(nextFeedingTime) {
		return fmt.Errorf("animal %d is not hungry yet", fs.AnimalID)
	}
	fs.LastFedTime = time.Now()
	return nil
}

func (fs *FeedingSchedule) ChangeFeedInterval(newInterval int) {
	fs.FeedInterval = newInterval
}
