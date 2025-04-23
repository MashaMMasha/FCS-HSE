package schedule_test

import (
	"testing"
	"time"
	"zoo/domain/schedule"

	"github.com/stretchr/testify/assert"
)

func TestNewFeedingSchedule(t *testing.T) {
	animalID := 1
	feedInterval := 6
	foodType := "Meat"

	fs := schedule.NewFeedingSchedule(animalID, feedInterval, foodType)

	assert.Equal(t, animalID, fs.AnimalID)
	assert.Equal(t, feedInterval, fs.FeedInterval)
	assert.Equal(t, foodType, fs.FoodType)
	assert.WithinDuration(t, time.Now(), fs.LastFedTime, time.Second)
}

func TestFeedAnimal_Success(t *testing.T) {
	fs := &schedule.FeedingSchedule{
		AnimalID:     1,
		FeedInterval: 6,
		LastFedTime:  time.Now().Add(-7 * time.Hour),
		FoodType:     "Meat",
	}

	err := fs.FeedAnimal()

	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), fs.LastFedTime, time.Second)
}

func TestFeedAnimal_NotHungry(t *testing.T) {
	fs := &schedule.FeedingSchedule{
		AnimalID:     1,
		FeedInterval: 6,
		LastFedTime:  time.Now().Add(-5 * time.Hour),
		FoodType:     "Meat",
	}

	err := fs.FeedAnimal()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is not hungry yet")
}

func TestChangeFeedInterval(t *testing.T) {
	fs := &schedule.FeedingSchedule{
		AnimalID:     1,
		FeedInterval: 6,
		LastFedTime:  time.Now(),
		FoodType:     "Meat",
	}

	newInterval := 8
	fs.ChangeFeedInterval(newInterval)

	assert.Equal(t, newInterval, fs.FeedInterval)
}
