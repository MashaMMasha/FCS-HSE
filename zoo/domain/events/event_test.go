package events_test

import (
	"testing"
	"zoo/domain/events"
)

func TestAnimalMovedEvent(t *testing.T) {
	event := events.NewAnimalMovedEvent(1, 2, 3)

	if event.AnimalID != 1 || event.OldEnclosureID != 2 || event.NewEnclosureID != 3 {
		t.Errorf("AnimalMovedEvent fields are not set correctly")
	}
}

func TestFeedingTimeEvent(t *testing.T) {
	event := events.NewFeedingTimeEvent(1, 1, "meat")

	if event.AnimalID != 1 || event.FoodType != "meat" {
		t.Errorf("FeedingEvent fields are not set correctly")
	}
}
